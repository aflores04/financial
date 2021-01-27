package service

import (
	"github.com/aflores04/financial/transaction/domain"
	"github.com/aflores04/financial/account/repository"
)

type IAccountService interface {
	Balance() float64
	PerformTransaction(transaction *domain.Transaction, success chan bool)
	Debit(amount float64)
	Credit(amount float64)
}

type AccountService struct {
	repository repository.IAccountRepository
}

func NewAccountService(repository repository.IAccountRepository) IAccountService {
	return &AccountService{
		repository: repository,
	}
}

func (s *AccountService) Balance() float64 {
	return s.repository.Balance()
}

func (s *AccountService) PerformTransaction(transaction *domain.Transaction, success chan bool) {
	if transaction.Type == domain.CREDIT {
		s.Credit(transaction.Amount)
	}

	if transaction.Type == domain.DEBIT {
		s.Debit(transaction.Amount)
	}

	if s.repository.Balance() < 0 {
		s.Credit(transaction.Amount)

		success <- false
	}

	success <- true
}

func (s *AccountService) Debit(amount float64) {
	s.repository.Debit(amount)
}

func (s *AccountService) Credit(amount float64) {
	s.repository.Credit(amount)
}