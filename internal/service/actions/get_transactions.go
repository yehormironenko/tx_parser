package actions

import (
	"github.com/yehormironenko/tx_parser/internal/model"
	"github.com/yehormironenko/tx_parser/internal/service"
	"github.com/yehormironenko/tx_parser/internal/service/helpers"
	"log"
)

type GetTransactionsService struct {
	externalClients service.ExternalClient
	logger          *log.Logger
}

func NewGetTransactionsService(clients service.ExternalClient, logger *log.Logger) *GetTransactionsService {
	return &GetTransactionsService{
		externalClients: clients,
		logger:          logger,
	}
}

func (gt *GetTransactionsService) GetTransactions(address string) (model.Transactions, error) {
	gt.logger.Printf("Getting all transactionf for adress: %v", address)

	resp, err := gt.externalClients.EthereumClient.GetTransactions(&address, nil, nil)
	if err != nil {
		gt.logger.Fatalf("Error response from externalclient: %v", err)
	}

	var transactions model.Transactions

	for _, result := range resp.Result {

		blockNumber, err := helpers.ConvertHexToInt(result.BlockNumber)
		if err != nil {
			gt.logger.Printf("Error converting blockNumber: %v\n", err)
		}
		transactionIndex, err := helpers.ConvertHexToInt(result.TransactionIndex)
		if err != nil {
			gt.logger.Printf("Error converting transactionIndex: %v\n", err)
		}
		logIndex, err := helpers.ConvertHexToInt(result.LogIndex)
		if err != nil {
			gt.logger.Printf("Error converting logIndex: %v\n", err)
		}

		transactions = append(transactions, model.Transaction{
			Address:          result.Address,
			Amount:           result.Data,
			BlockNumber:      int(blockNumber),
			LogIndex:         int(logIndex),
			TransactionIndex: int(transactionIndex),
			Removed:          result.Removed,
		})
	}
	gt.logger.Printf("Received transactions: %v", transactions)
	return transactions, nil
}
