package transaction

import (
	"fmt"
	"log"

	"github.com/jmoiron/sqlx"
	"github.com/lib/pq"
)

type Repository struct {
	DB *sqlx.DB
}

func (r *Repository) SaveTransaction(request TransactionRequest) error {
	queryInsertTransaction := `
    INSERT INTO tb_transactions (id, account_id, amount, currency, country, merchant, ip_address, occurred_at)
    VALUES ($1, $2, $3, $4, $5, $6, $7, $8)`

	log.Printf("INSERT tb_transactions: id=%s account=%s amount=%s currency=%s country=%s occurred_at=%s",
		request.TransactionID, request.AccountID, request.Amount, request.Currency, request.Country, request.OccurredAt)

	_, err := r.DB.Exec(queryInsertTransaction, request.TransactionID, request.AccountID, request.Amount, request.Currency, request.Country, request.Merchant, request.IPAddress, request.OccurredAt)

	if err != nil {
		log.Printf("failed to INSERT tb_transactions: %v", err)
		return fmt.Errorf("Error on insert Request %w", err)
	}

	return nil
}

func (r *Repository) SaveTransactionResponse(response TransactionResponse) error {
	queryInsertResponse := `
    INSERT INTO tb_scoring_results (id, transaction_id, score, risk, action, reasons)
    VALUES ($1, $2, $3, $4, $5, $6)`

	log.Printf("INSERT tb_scoring_results: id=%s transaction_id=%s score=%d risk=%s action=%s reasons=%v",
		response.ResponseID, response.TransactionID, response.Score, response.Risk, response.Action, response.Reasons)

	_, err := r.DB.Exec(queryInsertResponse, response.ResponseID, response.TransactionID, response.Score, response.Risk, response.Action, pq.Array(response.Reasons))

	if err != nil {
		log.Printf("failed to INSERT tb_scoring_results: %v", err)
		return fmt.Errorf("Error on insert Response %w", err)
	}

	return nil
}
