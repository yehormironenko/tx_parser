package actions

import (
	"fmt"
	"log"
	"sync"
	"time"

	ethlog "github.com/yehormironenko/tx_parser/internal/client/model"
	"github.com/yehormironenko/tx_parser/internal/model"
	"github.com/yehormironenko/tx_parser/internal/repository"
	"github.com/yehormironenko/tx_parser/internal/service"
	"github.com/yehormironenko/tx_parser/internal/service/helpers"
)

type NotificationService struct {
	mu              sync.Mutex
	repository      repository.SubscriberRepository
	externalClients service.ExternalClient
	currentBlock    int
	notifications   chan model.Transaction
	logger          *log.Logger
}

func NewNotificationService(repository repository.SubscriberRepository, client service.ExternalClient, logger *log.Logger) *NotificationService {
	return &NotificationService{
		repository:      repository,
		externalClients: client,
		notifications:   make(chan model.Transaction, 100),
		logger:          logger,
	}
}

// Poll the Ethereum blockchain for new blocks and transactions
func (e *NotificationService) StartPolling() {
	// Set current block
	currentBlock, err := e.externalClients.EthereumClient.GetCurrentBlock()
	if err != nil {
		e.logger.Println("Cannot check current log")
		return
	}
	currentHexBlock, err := helpers.ConvertHexToInt(currentBlock.Result)
	if err != nil {
		e.logger.Println("Cannot convert block to int")
		return
	}

	e.mu.Lock() // Lock to avoid race condition
	e.currentBlock = int(currentHexBlock)
	e.mu.Unlock()

	for {
		latestBlock, err := e.fetchLatestBlock()
		if err != nil {
			e.logger.Printf("Error fetching latest block: %v", err)
			time.Sleep(20 * time.Second) // Retry after a delay
			continue
		}

		// If we have new blocks, we want to check all transactions in the block, we need to wait for it to complete
		// Example: currentBlock =1 ; LatestBlock = 3 => we will check transactions for block 2
		if latestBlock > e.currentBlock+1 {
			for blockNum := e.currentBlock + 1; blockNum < latestBlock; blockNum++ {
				e.processBlock(latestBlock - 1)
			}
			e.currentBlock = latestBlock - 1
		}

		time.Sleep(20 * time.Second) // Poll every 20 seconds
	}
}

func (e *NotificationService) ProcessNotifications() {
	for transaction := range e.notifications {
		go func(tx model.Transaction) {
			e.logger.Printf("New transaction for address %s: %+v\n", tx.Address, tx)
			e.notifySubscriber(tx)
		}(transaction)
	}
}

// fetchLatestBlock retrieves the latest block from the Ethereum blockchain
func (e *NotificationService) fetchLatestBlock() (int, error) {
	e.logger.Println("Checking the latest block")
	jsonResponse, err := e.externalClients.EthereumClient.GetCurrentBlock()
	if err != nil {
		return 0, err
	}

	hexValue := jsonResponse.Result

	intValue, err := helpers.ConvertHexToInt(hexValue)
	if err != nil {
		return 0, fmt.Errorf("error converting hex to int: %w", err)
	}

	e.logger.Printf("Current block number is: %d", intValue)

	return int(intValue), nil
}

func (e *NotificationService) processBlock(blockNumber int) {
	toBlock := fmt.Sprintf("0x%x", blockNumber)

	// curl -X POST --data '{"jsonrpc":"2.0","method":"eth_getLogs","params":[{"fromBlock":"0x13cb930","toBlock":"latest"}],"id":1}' https://ethereum-rpc.publicnode.com/
	// Request above returns error: Please specify an address in your request
	for address, latestBlockNumber := range e.repository.GetSubscribers() {

		transactions, err := e.externalClients.EthereumClient.GetTransactions(&address, &latestBlockNumber, &toBlock)
		if err != nil {
			e.logger.Printf("Cannot get transactions for block %v: %v", blockNumber, err)
			return
		}

		if len(transactions.Result) > 0 {
			for _, v := range transactions.Result {
				transaction := e.convertLogToTransaction(v)

				go func(tx model.Transaction) {
					e.notifications <- tx
				}(transaction)
			}

		}
		_, err = e.repository.UpdateValue(address, toBlock)
		if err != nil {
			e.logger.Printf("Cannot update latestBlock for address %v: to %v", address, toBlock)
			return
		}
	}

}

// Example notification function (for WebSocket, push notifications, or emails)
func (e *NotificationService) notifySubscriber(tx model.Transaction) {
	e.logger.Printf("Notifying subscriber about transaction for address: %s", tx.Address)
	// TODO possible to configure different notifications
}

func (e *NotificationService) convertLogToTransaction(ethLog ethlog.EthLog) model.Transaction {
	blockNumber, err := helpers.ConvertHexToInt(ethLog.BlockNumber)
	if err != nil {
		e.logger.Printf("Error converting blockNumber: %v\n", err)
	}
	transactionIndex, err := helpers.ConvertHexToInt(ethLog.TransactionIndex)
	if err != nil {
		e.logger.Printf("Error converting transactionIndex: %v\n", err)
	}
	logIndex, err := helpers.ConvertHexToInt(ethLog.LogIndex)
	if err != nil {
		e.logger.Printf("Error converting logIndex: %v\n", err)
	}

	return model.Transaction{
		Address:          ethLog.Address,
		Amount:           ethLog.Data,
		BlockNumber:      int(blockNumber),
		LogIndex:         int(logIndex),
		TransactionIndex: int(transactionIndex),
		Removed:          ethLog.Removed,
	}
}
