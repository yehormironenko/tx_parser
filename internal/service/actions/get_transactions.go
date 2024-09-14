package actions

import (
	"log"
	"tx_parser/internal/model"
	"tx_parser/internal/service"
	"tx_parser/internal/service/helpers"
)

type GetTransactions struct {
	externalClients service.ExternalClient
	logger          *log.Logger
}

func NewGetTransactions(clients service.ExternalClient, logger *log.Logger) *GetTransactions {
	return &GetTransactions{
		externalClients: clients,
		logger:          logger,
	}
}

func (gt *GetTransactions) GetTransactions(address string) (model.Transactions, error) {

	log.Println("Request in GetTransactions")

	resp, err := gt.externalClients.EthereumClient.GetTransactions(address)
	if err != nil {
		log.Fatalf("Error response from externalclient: %v", err)
	}

	var transactions model.Transactions

	for _, result := range resp.Result {

		blockNumber, err := helpers.ConvertHexToInt(result.BlockNumber)
		if err != nil {
			log.Printf("Error converting blockNumber: %v\n", err)
		}
		transactionIndex, err := helpers.ConvertHexToInt(result.TransactionIndex)
		if err != nil {
			log.Printf("Error converting transactionIndex: %v\n", err)
		}
		logIndex, err := helpers.ConvertHexToInt(result.LogIndex)
		if err != nil {
			log.Printf("Error converting logIndex: %v\n", err)
		}

		transactions = append(transactions, model.Transaction{
			Address:          result.Address,
			Amount:           result.Data,
			BlockNumber:      blockNumber,
			LogIndex:         logIndex,
			TransactionIndex: transactionIndex,
			Removed:          result.Removed,
		})
	}

	return transactions, nil
}
