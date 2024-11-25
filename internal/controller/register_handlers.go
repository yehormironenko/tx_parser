package controller

import (
	"log"
	"net/http"

	"github.com/yehormironenko/tx_parser/internal/controller/handlers"
	"github.com/yehormironenko/tx_parser/internal/route"
	"github.com/yehormironenko/tx_parser/internal/service/actions"
	"github.com/yehormironenko/tx_parser/internal/service/core"
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
	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		http.NotFound(w, r)
	})
	return mux
}
