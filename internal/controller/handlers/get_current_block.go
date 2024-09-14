package handlers

import (
	"encoding/json"
	"net/http"
	"tx_parser/internal/service/actions"
)

func GetCurrentBlockHandler(service actions.GetCurrentBlock) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// Call the service to get the current block number
		blockNumber, err := service.GetCurrentBlock()
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		//TODO to model
		response := struct {
			BlockNumber int `json:"blockNumber"`
		}{
			BlockNumber: blockNumber,
		}

		// Set content type to application/json
		w.Header().Set("Content-Type", "application/json")

		if err = json.NewEncoder(w).Encode(response); err != nil {
			http.Error(w, "Error encoding response", http.StatusInternalServerError)
		}
	}
}