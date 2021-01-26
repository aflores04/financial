package service_test

import (
	"github.com/stretchr/testify/assert"
	"github.com/aflores04/financial/transaction/domain"
	"testing"
	"github.com/aflores04/financial/transaction/repository"
	"github.com/aflores04/financial/transaction/service"
)

func GetService() service.ITransactionService {
	repository := repository.NewTransactionRepository()
	return service.NewTransactionService(repository)
}

func TestGetHistory(t *testing.T) {
	var (
		service = GetService()
		expected []domain.Transaction
	)

	assert.Equal(t, expected, service.GetHistory())
}