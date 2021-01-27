package api

import (
	"github.com/aflores04/financial/transaction/handler"
	"github.com/aflores04/financial/transaction/service"
	"github.com/aflores04/financial/transaction/repository"
	accRepo "github.com/aflores04/financial/account/repository"
	accService "github.com/aflores04/financial/account/service"
	accHandler "github.com/aflores04/financial/account/handler"
	"github.com/gin-gonic/gin"
)

func InitRoutes() *gin.Engine {
	var (
		accountRepository		= accRepo.NewAccountRepository()
		accountService			= accService.NewAccountService(accountRepository)
		accountHandler			= accHandler.NewAccountHandler(accountService)
		transactionRepository 	= repository.NewTransactionRepository()
		transactionService 		= service.NewTransactionService(transactionRepository, accountService)		
		transactionHandler 		= handler.NewTransactionHandler(transactionService)

		router = gin.Default()
	)

	v1 := router.Group("/api/v1")

	transaction := v1.Group("transaction")
	transaction.GET("/", transactionHandler.HistoryHandler)
	transaction.POST("/", transactionHandler.CreateHandler)
	transaction.GET("/:id", transactionHandler.FindHandler)

	account := v1.Group("account")
	account.GET("/balance", accountHandler.BalanceHandler)

	return router
}