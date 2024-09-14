package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"tx_parser/config"
	"tx_parser/internal/client"
	"tx_parser/internal/controller"
	"tx_parser/internal/repository"
	"tx_parser/internal/service"
	"tx_parser/internal/service/actions"
	"tx_parser/internal/service/core"
)

func main() {
	logger := log.New(os.Stdout, "CONFIG-LOADER: ", log.Ldate|log.Ltime|log.Lshortfile)

	c, err := config.LoadConfig("config/config.json", logger)
	if err != nil {
		log.Panicf("cannot read config file")
	}
	logger.SetPrefix("")
	//externalClient
	ethereumClient := client.NewEthereumApiClient(c.Client.Endpoint, logger)

	subscriberMap := make(map[string]struct{})
	inMemoryRepository := repository.NewMemoryRepo(subscriberMap, logger)

	//TODO create a builder for it
	// services
	echoService := core.NewEcho(logger)
	getCurrentBlockNumberService := actions.NewGetCurrentBlockService(service.ExternalClient{EthereumClient: ethereumClient}, logger)
	getTransactionsService := actions.NewGetTransactionsService(service.ExternalClient{EthereumClient: ethereumClient}, logger)
	notificationService := actions.NewNotificationService(inMemoryRepository, logger)
	handlerSettings := &controller.HandlersSettings{
		EchoService:     echoService,
		GetCurrentBlock: getCurrentBlockNumberService,
		GetTransactions: getTransactionsService,
		Subscriptions:   notificationService,
		Logger:          logger,
	}

	mux := controller.Handlers(handlerSettings)

	log.Printf("Server started on %v:%v", c.Server.Host, c.Server.Port)
	err = http.ListenAndServe(fmt.Sprintf("%v:%v", c.Server.Host, c.Server.Port), mux)
	if err != nil {
		log.Fatalf("Server failed to start: %v", err)
	}

}
