package transaction

import (
	"log"
	"strings"
	"time"

	"github.com/google/uuid"
	"github.com/shopspring/decimal"
)

type EngineService struct {
	Repo *Repository
}

func (s *EngineService) Evaluate(transaction TransactionRequest) (TransactionResponse, error) {

	transaction.TransactionID = uuid.New()
	transaction.OccurredAt = time.Now().UTC()

	if err := s.Repo.SaveTransaction(transaction); err != nil {
		log.Printf("failed to save transaction: %v", err)
		return TransactionResponse{}, err
	}
	log.Printf("transaction saved: transaction_id=%s", transaction.TransactionID)

	highValuePoints, highValueReason := highValueRule(transaction)

	nightTransactionPoints, nightTransactionReason := nightTimeRule(transaction)

	foreignCountryPoints, foreignCountryReason := foreignCountryRule(transaction)

	suspiciousMerchantPoints, suspeciousMerchantReason := suspiciousMerchantRule(transaction)

	roundNumberPoints, roundNumberReason := roundNumberRule(transaction)

	score := sumScore(highValuePoints, nightTransactionPoints, foreignCountryPoints, suspiciousMerchantPoints, roundNumberPoints)

	reasonList := []string{}

	if highValuePoints != 0 {
		reasonList = append(reasonList, highValueReason)
	}

	if nightTransactionPoints != 0 {
		reasonList = append(reasonList, nightTransactionReason)
	}

	if foreignCountryPoints != 0 {
		reasonList = append(reasonList, foreignCountryReason)
	}

	if suspiciousMerchantPoints != 0 {
		reasonList = append(reasonList, suspeciousMerchantReason)
	}

	if roundNumberPoints != 0 {
		reasonList = append(reasonList, roundNumberReason)
	}

	if score > 100 {
		score = 100
	}

	risk, action := classifier(score)

	transactionResponse := TransactionResponse{
		ResponseID:    uuid.New(),
		TransactionID: transaction.TransactionID,
		Score:         score,
		Risk:          risk,
		Action:        action,
		Reasons:       reasonList,
	}

	if err := s.Repo.SaveTransactionResponse(transactionResponse); err != nil {
		log.Printf("failed to save response: %v", err)
		return TransactionResponse{}, err
	}
	log.Printf("response saved: response_id=%s", transactionResponse.ResponseID)

	return transactionResponse, nil
}

func highValueRule(transaction TransactionRequest) (int, string) {

	limit := decimal.NewFromInt(10000)
	if transaction.Amount.GreaterThan(limit) {
		return 30, "high_value"

	}

	return 0, ""
}

func foreignCountryRule(transaction TransactionRequest) (int, string) {
	if transaction.Country != "BR" {
		return 40, "foreign_country"
	}
	return 0, ""
}

func suspiciousMerchantRule(transaction TransactionRequest) (int, string) {
	casino := strings.Contains(strings.ToLower(transaction.Merchant), "casino")

	crypto := strings.Contains(strings.ToLower(transaction.Merchant), "crypto")

	bet := strings.Contains(strings.ToLower(transaction.Merchant), "bet")

	if casino || crypto || bet {
		return 35, "suspicious_merchant"
	}

	return 0, ""
}

func nightTimeRule(transaction TransactionRequest) (int, string) {
	loc, _ := time.LoadLocation("America/Sao_Paulo")
	occurredAtBR := transaction.OccurredAt.In(loc)

	if occurredAtBR.Hour() >= 0 && occurredAtBR.Hour() < 5 {
		return 25, "night_transaction"
	}

	return 0, ""
}

func roundNumberRule(transaction TransactionRequest) (int, string) {

	if transaction.Amount.Mod(decimal.NewFromInt(1000)).Equal(decimal.Zero) {
		return 10, "round_number"
	}

	return 0, ""
}

func sumScore(values ...int) int { // Funnção variádica no Go

	score := 0

	for _, v := range values {
		score += v
	}

	return score

}

func classifier(score int) (string, string) {
	switch {
	case score <= 30:
		return "low", "approve"
	case score <= 60:
		return "medium", "review"
	case score <= 85:
		return "high", "review"
	default:
		return "critical", "block"
	}
}
