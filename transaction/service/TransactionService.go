package service

import (
	"errors"
	"github.com/aflores04/financial/account/service"
	"github.com/google/uuid"
	"github.com/aflores04/financial/transaction/request"
	"github.com/aflores04/financial/transaction/repository"
	"github.com/aflores04/financial/transaction/domain"
)

type ITransactionService interface {
	GetHistory() []domain.Transaction
	Create(request request.CreateTransaction) (*domain.Transaction, error)
	Find(id uuid.UUID) domain.Transaction
}

type TransactionService struct {
	repository repository.ITransactionRepository
	accountService service.IAccountService
}

func NewTransactionService(
	repository repository.ITransactionRepository,
	accountService service.IAccountService) ITransactionService {
	return &TransactionService{
		repository: repository,
		accountService: accountService,
	}
}

func (s TransactionService) GetHistory() []domain.Transaction {
	return s.repository.GetHistory()
}

func (s TransactionService) Create(request request.CreateTransaction) (*domain.Transaction, error) {
	var (
		transaction *domain.Transaction
		successChannel = make(chan bool)
		transactionChannel = make(chan *domain.Transaction)
	)

	go s.repository.Create(request.Amount, request.Type, transactionChannel)

	transaction = <- transactionChannel
	
	go s.accountService.PerformTransaction(transaction, successChannel)	

	success := <- successChannel

	if !success {
		return transaction, errors.New("invalid transaction")
	}

	return transaction, nil
}

func (s TransactionService) Find(id uuid.UUID) domain.Transaction {
	return s.repository.Find(id)
}