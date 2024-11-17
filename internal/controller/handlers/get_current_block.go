package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/yehormironenko/tx_parser/internal/model"
	"github.com/yehormironenko/tx_parser/internal/service"
)

func GetCurrentBlockHandler(service service.BlockRetriever) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// Call the service to get the current block number
		blockNumber, err := service.GetCurrentBlock()
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		response := model.GetCurrentBlockResponse{BlockNumber: blockNumber}

		// Set content type to application/json
		w.Header().Set("Content-Type", "application/json")

		if err = json.NewEncoder(w).Encode(response); err != nil {
			http.Error(w, "Error encoding response", http.StatusInternalServerError)
		}
	}
}
