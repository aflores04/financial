package repository

import (
	"log"
	"time"
	"github.com/google/uuid"
	"github.com/aflores04/financial/transaction/domain"
)

type ITransactionRepository interface {
	GetHistory() []domain.Transaction
	Create(amount float64, transactionType string, transactionChannel chan *domain.Transaction)
	Find(id uuid.UUID) domain.Transaction 
}

type TransactionRepository struct {
	History []domain.Transaction
}

func NewTransactionRepository() ITransactionRepository {
	return &TransactionRepository{}
}

func (r *TransactionRepository) Create(
	amount float64, 
	transactionType string, 
	transactionChannel chan *domain.Transaction)  {
	currentTime := time.Now()

	transaction := domain.Transaction{
		ID: uuid.New(),
		Type: transactionType,
		Amount: amount,
		EffectiveDate: currentTime.Format("2006.01.02 15:04:05"),
	}

	r.History = append(r.History, transaction)

	transactionChannel <- &transaction
}

func (r *TransactionRepository) GetHistory() []domain.Transaction {
	return r.History
}

func (r *TransactionRepository) Find(id uuid.UUID) domain.Transaction {
	log.Println(id)
	for _, transaction := range r.History {
		if transaction.ID == id {
			return transaction
		}
	}

	panic("transaction not found")
}