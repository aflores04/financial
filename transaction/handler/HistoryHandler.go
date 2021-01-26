package handler

import (
	"log"
	"net/http"
	"github.com/gin-gonic/gin"
)

func (h *TransactionHandler) HistoryHandler(c *gin.Context) {
	c.JSON(http.StatusOK, h.service.GetHistory())
}
