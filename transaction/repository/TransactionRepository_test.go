package repository_test

import (
	"financial/transaction/repository"
	"financial/transaction/domain"
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

