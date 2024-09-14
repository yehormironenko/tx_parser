package handlers

import (
	"net/http"
	"tx_parser/internal/service/core"
)

func EchoHandler(service core.Echo) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		echoResponse := service.Echo()
		w.Write([]byte(echoResponse))
	}
}
