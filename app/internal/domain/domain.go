package domain

import (
	"time"

	"github.com/google/uuid"
	"github.com/shopspring/decimal"
)

type TransactionRequest struct {
	TransactionID uuid.UUID       `db:"id" json:"transaction_id"`
	AccountID     string          `db:"account_id" json:"account_id"`
	Amount        decimal.Decimal `db:"amount" json:"amount"`
	Currency      string          `db:"currency" json:"currency"`
	Country       string          `db:"country" json:"country"`
	Merchant      string          `db:"merchant" json:"merchant"`
	IPAddress     string          `db:"ip_address" json:"ip_address"`
	OccurredAt    time.Time       `db:"occurred_at" json:"occurred_at"`
}

type TransactionResponse struct {
	ResponseID    uuid.UUID `db:"id" json:"response_id"`
	TransactionID uuid.UUID `db:"transaction_id" json:"transaction_id"`
	Score         int       `db:"score" json:"score"`
	Risk          string    `db:"risk" json:"risk"`
	Action        string    `db:"action" json:"action"`
	Reasons       []string  `db:"reasons" json:"reasons"`
}
