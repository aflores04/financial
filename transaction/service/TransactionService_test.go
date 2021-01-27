package service_test

import (
	"github.com/google/uuid"
	"github.com/aflores04/financial/transaction/request"
	"time"
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
	currentTime := time.Now()
	var (
		service = GetService()
		request = request.CreateTransaction{
			Amount: 5,
			Type: "credit",
		}
	)

	transaction := service.Create(request)

	expected := []domain.Transaction{
		{
			ID: transaction.ID,
			Amount: 5,
			Type: "credit",
			EffectiveDate: currentTime.Format("2006.01.02 15:04:05"),
		},
	}

	assert.Equal(t, expected, service.GetHistory())
}

func TestCreate(t *testing.T) {
	currentTime := time.Now()
	var (
		service = GetService()
		request = request.CreateTransaction{
			Amount: 5,
			Type: "credit",
		}
	)

	transaction := service.Create(request)

	expected := domain.Transaction{
		ID: transaction.ID,
		Amount: 5,
		Type: "credit",
		EffectiveDate: currentTime.Format("2006.01.02 15:04:05"),
	}

	assert.Equal(t, expected, transaction)
}

func TestFindTransactionPanic(t *testing.T) {
	var (
		service = GetService()
	)

	_, err := service.Find(uuid.New())

	assert.Equal(t, err, "transaction not found")
}