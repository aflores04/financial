package repository

import (
	"github.com/aflores04/financial/transaction/domain"
)

type ITransactionRepository interface {
	GetHistory() []domain.Transaction
}

type TransactionRepository struct {
	History []domain.Transaction
}

func NewTransactionRepository() ITransactionRepository {
	return &TransactionRepository{}
}

func (r *TransactionRepository) GetHistory() []domain.Transaction {
	return r.History
}