package transaction

import (
	"encoding/json"
	"net/http"
)

type TransactionHandler struct {
	Service *EngineService
}

func (h *TransactionHandler) Transaction(w http.ResponseWriter, r *http.Request) {
	var transaction TransactionRequest
	if err := json.NewDecoder(r.Body).Decode(&transaction); err != nil {
		http.Error(w, "invalid request body", http.StatusBadRequest)
		return
	}

	if err := h.Service.SaveTransaction(transaction); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}
