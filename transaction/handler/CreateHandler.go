package handler

import (
	"log"
	"github.com/aflores04/financial/transaction/response"
	"github.com/aflores04/financial/transaction/request"
	"net/http"
	"github.com/gin-gonic/gin"
)

func (h *TransactionHandler) CreateHandler(c *gin.Context) {
	var postRequest request.CreateTransaction

	if err := c.BindJSON(&postRequest); err != nil {
		c.JSON(http.StatusBadRequest, response.ErrorResponse{
			Code:		http.StatusBadRequest,
			Message:	"invalid input",
		})
		return
	}

	log.Println(postRequest.Amount)

	c.JSON(http.StatusCreated, h.service.Create(postRequest))
}
