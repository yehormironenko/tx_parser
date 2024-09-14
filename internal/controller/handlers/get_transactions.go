package handlers

import (
	"encoding/json"
	"net/http"
	"tx_parser/internal/model"
	"tx_parser/internal/service"
)

func GetTransactionsHandler(service service.TransactionFetcher) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req model.GetTransactionsRequest
		if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
			http.Error(w, "Invalid request body", http.StatusBadRequest)
			return
		}

		if req.Address == "" {
			http.Error(w, "Address is required", http.StatusBadRequest)
			return
		}

		transactions, err := service.GetTransactions(req.Address)
		if err != nil {
			http.Error(w, "Error retrieving transactions", http.StatusInternalServerError)
			return
		}

		// Prepare and send the response
		w.Header().Set("Content-Type", "application/json")
		if err = json.NewEncoder(w).Encode(transactions); err != nil {
			http.Error(w, "Error encoding response", http.StatusInternalServerError)
		}
	}
}
