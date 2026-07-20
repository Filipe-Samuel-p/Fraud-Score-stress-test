package transaction

type EngineService struct {
	Repo *Repository
}

func (s *EngineService) SaveTransaction(transaction TransactionRequest) error {
	err := s.Repo.SaveTransaction(transaction)
	if err != nil {
		return err
	}

	return nil
}
