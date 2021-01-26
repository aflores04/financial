package api

import (
	"github.com/aflores04/financial/transaction/handler"
	"github.com/aflores04/financial/transaction/service"
	"github.com/aflores04/financial/transaction/repository"
	"github.com/gin-gonic/gin"
)

func InitRoutes() *gin.Engine {
	var (
		TransactionRepository 	= repository.NewTransactionRepository()
		TransactionService 		= service.NewTransactionService(TransactionRepository)		
		TransactionHandler 		= handler.NewTransactionHandler(TransactionService)

		router = gin.Default()
	)

	v1 := router.Group("/api/v1")

	transaction := v1.Group("transaction")
	transaction.GET("/", TransactionHandler.HistoryHandler)

	return router
}