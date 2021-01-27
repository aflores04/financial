package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/aflores04/financial/account/service"
)

type IAccountHandler interface {
	BalanceHandler(c *gin.Context)
}

type AccountHandler struct {
	service service.IAccountService
}

func NewAccountHandler(service service.IAccountService) IAccountHandler {
	return &AccountHandler{
		service: service,
	}
}