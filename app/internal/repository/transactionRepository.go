package repository

import (
	"fmt"
	"fraud-score/internal/domain"
	"time"

	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
)

type Repository struct {
	DB *sqlx.DB
}

func (r *Repository) SaveTransaction(request domain.TransactionRequest) error {
	request.TransactionID = uuid.New()
	request.OccurredAt = time.Now()

	queryInsertTransaction := `
    INSERT INTO tb_transactions (id, account_id, amount, currency, country, merchant, ip_address, occurred_at)
    VALUES ($1, $2, $3, $4, $5, $6, $7, $8)`

	_, err := r.DB.Exec(queryInsertTransaction, request.TransactionID, request.AccountID, request.Amount, request.Currency, request.Country, request.Merchant, request.IPAddress, request.OccurredAt)

	if err != nil {
		return fmt.Errorf("Error on insert Request %w", err)
	}

	return nil
}
