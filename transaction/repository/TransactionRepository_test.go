package repository_test

import (
	"github.com/google/uuid"
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
	currentTime := time.Now()
	var (
		repository = GetRepository()
		expected []domain.Transaction
	)

	transaction := repository.Create(10, "credit")
	expected = []domain.Transaction{
		domain.Transaction{
			ID: transaction.ID,
			Amount: 10,
			Type: "credit",
			EffectiveDate: currentTime.Format("2006.01.02 15:04:05"),
		},
	}
	
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

func TestFindTransactionSuccess(t *testing.T) {
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

	assert.Equal(t, expected, repository.Find(transaction.ID))
}

func TestFindTransactionPanic(t *testing.T) {
	var (
		repository = GetRepository()
	)

	defer func() { 
		if r := recover(); r != nil {
			assert.Equal(t, r, "transaction not found")
		}
	}()

	repository.Find(uuid.New())
}