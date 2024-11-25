package actions

import (
	"fmt"
	"log"

	"github.com/yehormironenko/tx_parser/internal/service"
	"github.com/yehormironenko/tx_parser/internal/service/helpers"
)

type GetCurrentBlockService struct {
	externalClients service.ExternalClient
	logger          *log.Logger
}

func NewGetCurrentBlockService(clients service.ExternalClient, logger *log.Logger) *GetCurrentBlockService {
	return &GetCurrentBlockService{
		externalClients: clients,
		logger:          logger,
	}
}

func (gb *GetCurrentBlockService) GetCurrentBlock() (int, error) {
	gb.logger.Println("Checking current block number service")
	jsonResponse, err := gb.externalClients.EthereumClient.GetCurrentBlock()
	if err != nil {
		return 0, err
	}

	hexValue := jsonResponse.Result

	intValue, err := helpers.ConvertHexToInt(hexValue)
	if err != nil {
		return 0, fmt.Errorf("error converting hex to int: %w", err)
	}

	gb.logger.Printf("Current block number is: %d", intValue)

	return int(intValue), nil
}
