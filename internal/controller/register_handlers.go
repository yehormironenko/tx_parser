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
	GetCurrentBlock *actions.GetCurrentBlock
	Logger          *log.Logger
}

func Handlers(settings *HandlersSettings) *http.ServeMux {
	mux := http.NewServeMux()
	// Register the echo for "/echo"
	mux.HandleFunc(route.EchoPath, handlers.EchoHandler(*settings.EchoService))
	mux.HandleFunc(route.GetCurrentBlockPath, handlers.GetCurrentBlockHandler(*settings.GetCurrentBlock))
	//...all other handlers

	return mux
}
