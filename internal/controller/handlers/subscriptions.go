package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/yehormironenko/tx_parser/internal"
	"github.com/yehormironenko/tx_parser/internal/model"
	"github.com/yehormironenko/tx_parser/internal/service"
)

func SubscribeHandler(service service.SubscriptionManager) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req model.SubscribeRequest
		if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
			http.Error(w, "Invalid request body", http.StatusBadRequest)
			return
		}

		// Validate the action and address fields
		if req.Action == "" || req.Address == "" {
			http.Error(w, "Missing required fields", http.StatusBadRequest)
			return
		}

		switch req.Action {
		case internal.SubscribeActions:
			subscribed, err := service.Subscribe(req.Address)
			if err != nil {
				w.WriteHeader(http.StatusInternalServerError)
			} else if subscribed {
				w.WriteHeader(http.StatusOK)
				w.Write([]byte(fmt.Sprintf("Subscribed to address %s", req.Address)))
			} else {
				w.WriteHeader(http.StatusConflict)
				w.Write([]byte(fmt.Sprintf("Already subscribed to address %s", req.Address)))
			}

		case internal.UnsubscribeActions:
			unsubscribed, err := service.Unsubscribe(req.Address)
			if err != nil {
				w.WriteHeader(http.StatusInternalServerError)
			} else if unsubscribed {
				w.WriteHeader(http.StatusOK)
				w.Write([]byte(fmt.Sprintf("Unsubscribed from address %s", req.Address)))
			} else {
				w.WriteHeader(http.StatusNotFound)
				w.Write([]byte(fmt.Sprintf("Address %s was not subscribed", req.Address)))
			}

		default:
			http.Error(w, "Invalid action. Use 'subscribe' or 'unsubscribe'", http.StatusBadRequest)
		}
	}
}
