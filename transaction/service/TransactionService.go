package service

import (
	"financial/transaction/repository"
	"financial/transaction/domain"
)

type ITransactionService interface {
	GetHistory() []domain.Transaction
}

type TransactionService struct {
	repository repository.ITransactionRepository
}

func NewTransactionService(repository repository.ITransactionRepository) ITransactionService {
	return &TransactionService{
		repository: repository,
	}
}

func (s TransactionService) GetHistory() []domain.Transaction {
	return s.repository.GetHistory()
}
