package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"tx_parser/config"
	"tx_parser/internal/controller"
	"tx_parser/internal/service/core"
)

func main() {
	logger := log.New(os.Stdout, "CONFIG-LOADER: ", log.Ldate|log.Ltime|log.Lshortfile)

	c, err := config.LoadConfig("config/config.json", logger)
	if err != nil {
		log.Panicf("cannot read config file")
	}

	echoService := core.NewEcho(logger)
	handlerSettings := &controller.HandlersSettings{
		Logger:      logger,
		EchoService: echoService,
	}
	mux := controller.Handlers(handlerSettings)

	log.Printf("Server started on %v:%v", c.Server.Host, c.Server.Port)
	err = http.ListenAndServe(fmt.Sprintf("%v:%v", c.Server.Host, c.Server.Port), mux)
	if err != nil {
		log.Fatalf("Server failed to start: %v", err)
	}

}
