package handlers

import (
	"net/http"
	"tx_parser/internal/service"
)

func EchoHandler(service service.Echoer) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		echoResponse := service.Echo()
		w.Write([]byte(echoResponse))
	}
}
