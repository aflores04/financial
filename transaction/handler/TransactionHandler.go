package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/aflores04/financial/transaction/service"
)

type ITransactionHandler interface {
	HistoryHandler(c *gin.Context)
	CreateHandler(c *gin.Context)
	FindHandler(c *gin.Context)
}

type TransactionHandler struct {
	service service.ITransactionService
}

func NewTransactionHandler(service service.ITransactionService) ITransactionHandler {
	return &TransactionHandler{
		service: service,
	}
}