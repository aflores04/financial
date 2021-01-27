package repository

import (
	"github.com/aflores04/financial/account/domain"
)

type IAccountRepository interface {
	Debit(amount float64)
	Credit(amount float64)
	Balance() float64
}

type AccountRepository struct {
	account *domain.Account
}

func NewAccountRepository() IAccountRepository {
	return &AccountRepository{
		account: &domain.Account{
			Balance: 0,
		},
	}
}

func (r *AccountRepository) Balance() float64 {
	return r.account.Balance
}

func (r *AccountRepository) Debit(amount float64) {
	r.account.Balance = r.account.Balance - amount
}

func (r *AccountRepository) Credit(amount float64) {
	r.account.Balance = r.account.Balance + amount
}