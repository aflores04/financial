package service

import (
	"github.com/aflores04/financial/transaction/request"
	"github.com/aflores04/financial/transaction/repository"
	"github.com/aflores04/financial/transaction/domain"
)

type ITransactionService interface {
	GetHistory() []domain.Transaction
	Create(request request.CreateTransaction) domain.Transaction
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

func (s TransactionService) Create(request request.CreateTransaction) domain.Transaction {
	return s.repository.Create(request.Amount, request.Type)
}