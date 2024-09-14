package actions

import (
	"log"
	"tx_parser/internal/service"
)

type GetCurrentBlock struct {
	externalClients service.ExternalClient
	logger          *log.Logger
}

func NewGetCurrentBlock(clients service.ExternalClient, logger *log.Logger) *GetCurrentBlock {
	return &GetCurrentBlock{
		externalClients: clients,
		logger:          logger,
	}
}

func (gb *GetCurrentBlock) GetCurrentBlock() (int, error) {
	log.Println("Request in GetCurrentBlock")
	blockNumber, err := gb.externalClients.EthereumClient.GetCurrentBlock()
	if err != nil {
		return 0, err
	}
	log.Printf("Current block number is: %d", blockNumber)
	return blockNumber, nil
}
