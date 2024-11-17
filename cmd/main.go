package main

import (
	"context"
	"errors"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/yehormironenko/tx_parser/cmd/initialize"
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
	server := &http.Server{
		Addr:    fmt.Sprintf("%v:%v", components.Config.Server.Host, components.Config.Server.Port),
		Handler: components.Mux,
	}

	go func() {
		components.Logger.Printf("Server started on %v:%v", components.Config.Server.Host, components.Config.Server.Port)
		if err = server.ListenAndServe(); err != nil && !errors.Is(http.ErrServerClosed, err) {
			components.Logger.Fatalf("Server failed to start: %v", err)
		}
	}()

	gracefulShutdown(server, components.Logger)
}

func gracefulShutdown(server *http.Server, logger *log.Logger) {
	stopChan := make(chan os.Signal, 1)
	signal.Notify(stopChan, syscall.SIGINT, syscall.SIGTERM)

	// Block until a signal is received
	<-stopChan
	logger.Println("Shutting down server...")

	// Create a context with a timeout for the shutdown
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	// Attempt graceful shutdown
	if err := server.Shutdown(ctx); err != nil {
		logger.Fatalf("Server forced to shutdown: %v", err)
	}

	logger.Println("Server gracefully stopped")
}
