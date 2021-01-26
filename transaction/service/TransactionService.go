package service

import (
	"github.com/aflores04/financial/transaction/repository"
	"github.com/aflores04/financial/transaction/domain"
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
