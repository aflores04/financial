package repository_test

import (
	"time"
	"github.com/aflores04/financial/transaction/repository"
	"github.com/aflores04/financial/transaction/domain"
	"github.com/stretchr/testify/assert"
	"testing"
)

func GetRepository() repository.ITransactionRepository {
	return repository.NewTransactionRepository()
}

func TestGetHistorySuccess(t *testing.T) {
	var (
		repository = GetRepository()
		expected []domain.Transaction
	)
	
	assert.Equal(t, expected, repository.GetHistory())
} 

func TestCreateTransactionSuccess(t *testing.T) {
	currentTime := time.Now()
	var (
		repository = GetRepository()
		expected domain.Transaction
	)

	transaction := repository.Create(10, "credit")
	expected = domain.Transaction{
		ID: transaction.ID,
		Amount: 10,
		Type: "credit",
		EffectiveDate: currentTime.Format("2006.01.02 15:04:05"),
	}

	assert.NotEmpty(t, transaction)
	assert.Equal(t, expected, transaction)
}

