package main

import (
	"fmt"
	"log"
	"net/http"
	"tx_parser/cmd/initialize"
)

func main() {
	components, err := initialize.NewAppComponents()
	if err != nil {
		log.Fatalf("Error initializing application components: %v", err)
	}

	// Start polling Ethereum blockchain
	go components.NotificationService.StartPolling()
	// Start processing notifications asynchronously
	go components.NotificationService.ProcessNotifications()

	// Start HTTP server
	err = http.ListenAndServe(fmt.Sprintf("%v:%v", components.Config.Server.Host, components.Config.Server.Port), components.Mux)
	if err != nil {
		components.Logger.Fatalf("Server failed to start: %v", err)
	}
	components.Logger.Printf("Server started on %v:%v", components.Config.Server.Host, components.Config.Server.Port)
}
