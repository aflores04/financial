package service_test

import (
	"financial/transaction/repository"
	"financial/transaction/service"
)

func GetService() service.ITransactionService {
	repository = repository.NewTransactionRepository()
	return service.NewTransactionService(repository)
}