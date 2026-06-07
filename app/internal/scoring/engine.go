package scoring

import (
	"fraud-score/internal/domain"
	"fraud-score/internal/repository"
)

type EngineService struct {
	Repo *repository.Repository
}

func (s *EngineService) SaveTransaction(transaction domain.TransactionRequest) error {
	err := s.Repo.SaveTransaction(transaction)
	if err != nil {
		return err
	}

	return nil
}
