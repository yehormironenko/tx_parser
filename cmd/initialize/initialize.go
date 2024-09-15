package initialize

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

type AppComponents struct {
	Config                 *config.Config
	EchoService            *core.Echo
	GetCurrentBlockService *actions.GetCurrentBlockService
	GetTransactionsService *actions.GetTransactionsService
	SubscriptionService    *actions.SubscriptionsService
	NotificationService    *actions.NotificationService
	Logger                 *log.Logger
	Mux                    http.Handler
}

func NewAppComponents() (*AppComponents, error) {
	logger := log.New(os.Stdout, "CONFIG-LOADER: ", log.Ldate|log.Ltime|log.Lshortfile)

	c, err := config.LoadConfig("config/config.json", logger)
	if err != nil {
		return nil, fmt.Errorf("cannot read config file: %w", err)
	}
	logger.SetPrefix("")

	ethereumClient := client.NewEthereumApiClient(c.Client.Endpoint, logger)
	subscriberMap := make(map[string]string)
	inMemoryRepository := repository.NewMemoryRepo(subscriberMap, logger)

	echoService := core.NewEcho(logger)
	getCurrentBlockNumberService := actions.NewGetCurrentBlockService(service.ExternalClient{EthereumClient: ethereumClient}, logger)
	getTransactionsService := actions.NewGetTransactionsService(service.ExternalClient{EthereumClient: ethereumClient}, logger)
	subscriptionService := actions.NewSubscriptionsService(inMemoryRepository, service.ExternalClient{EthereumClient: ethereumClient}, logger)
	notificationService := actions.NewNotificationService(inMemoryRepository, service.ExternalClient{EthereumClient: ethereumClient}, logger)

	handlerSettings := &controller.HandlersSettings{
		EchoService:     echoService,
		GetCurrentBlock: getCurrentBlockNumberService,
		GetTransactions: getTransactionsService,
		Subscriptions:   subscriptionService,
		Logger:          logger,
	}
	mux := controller.Handlers(handlerSettings)

	return &AppComponents{
		Config:                 c,
		EchoService:            echoService,
		GetCurrentBlockService: getCurrentBlockNumberService,
		GetTransactionsService: getTransactionsService,
		SubscriptionService:    subscriptionService,
		NotificationService:    notificationService,
		Logger:                 logger,
		Mux:                    mux,
	}, nil
}
