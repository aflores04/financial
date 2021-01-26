package handler

import (
	"financial/transaction/service"
)

type ITransactionHandler interface {
	HistoryHandler(c *gin.Context)
}

type TransactionHandler struct {
	Service service.ITransactionService
}

func NewTransactionHandler() ITransactionHandler {
	return &TransactionHandler
}