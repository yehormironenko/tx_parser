package handlers

import (
	"github.com/yehormironenko/tx_parser/internal/service"
	"net/http"
)

func EchoHandler(service service.Echoer) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		echoResponse := service.Echo()
		w.Write([]byte(echoResponse))
	}
}
