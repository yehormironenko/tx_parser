package controller

import (
	"log"
	"net/http"
	"tx_parser/internal/controller/handlers"
	"tx_parser/internal/route"
	"tx_parser/internal/service/actions"
	"tx_parser/internal/service/core"
)

type HandlersSettings struct {
	EchoService     *core.Echo
	GetCurrentBlock *actions.GetCurrentBlockService
	GetTransactions *actions.GetTransactionsService
	Subscriptions   *actions.SubscriptionsService
	Logger          *log.Logger
}

func Handlers(settings *HandlersSettings) *http.ServeMux {
	mux := http.NewServeMux()
	mux.HandleFunc(route.EchoPath, handlers.EchoHandler(settings.EchoService))
	mux.HandleFunc(route.GetCurrentBlockPath, handlers.GetCurrentBlockHandler(settings.GetCurrentBlock))
	mux.HandleFunc(route.GetTransactionsPath, handlers.GetTransactionsHandler(settings.GetTransactions))
	mux.HandleFunc(route.SubscribePath, handlers.SubscribeHandler(settings.Subscriptions))

	//TODO add all other route

	return mux
}
