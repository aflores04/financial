package handler

import (
	"net/http"
	"github.com/gin-gonic/gin"
)

func (h *AccountHandler) BalanceHandler(c *gin.Context) {
	c.JSON(http.StatusOK, struct{
		Balance float64 `json:"balance"`
	}{
		Balance: h.service.Balance(),
	})
}