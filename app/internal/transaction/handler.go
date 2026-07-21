package transaction

import (
	"encoding/json"
	"log"
	"net/http"
)

type TransactionHandler struct {
	Service *EngineService
}

func (h *TransactionHandler) Transaction(w http.ResponseWriter, r *http.Request) {

	var transaction TransactionRequest
	if err := json.NewDecoder(r.Body).Decode(&transaction); err != nil {
		log.Printf("failed to decode body: %v", err)
		http.Error(w, "invalid request body", http.StatusBadRequest)
		return
	}

	response, err := h.Service.Evaluate(transaction)
	if err != nil {
		log.Printf("failed to evaluate transaction: %v", err)
		http.Error(w, "could not process transaction", http.StatusInternalServerError)
		return
	}

	log.Printf("transaction evaluated successfully: response_id=%s score=%d action=%s",
		response.ResponseID, response.Score, response.Action)

	w.WriteHeader(http.StatusCreated)
	if err := json.NewEncoder(w).Encode(response); err != nil {
		log.Printf("failed to encode response: %v", err)
	}
}
