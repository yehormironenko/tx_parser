package actions

import (
	"fmt"
	"log"
	"tx_parser/internal/service"
	"tx_parser/internal/service/helpers"
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
	jsonResponse, err := gb.externalClients.EthereumClient.GetCurrentBlock()
	if err != nil {
		return 0, err
	}

	hexValue := jsonResponse.Result

	intValue, err := helpers.ConvertHexToInt(hexValue)
	if err != nil {
		return 0, fmt.Errorf("error converting hex to int: %w", err)
	}

	log.Printf("Current block number is: %d", intValue)

	return int(intValue), nil
}
