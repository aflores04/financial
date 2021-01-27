package handler

import (
	"github.com/google/uuid"
	"github.com/aflores04/financial/transaction/response"
	"net/http"
	"github.com/gin-gonic/gin"
)

func (h *TransactionHandler) FindHandler(c *gin.Context) {
	idParam := c.Param("id")

	defer func() { 
		if r := recover(); r != nil {
			c.JSON(http.StatusBadRequest, response.ErrorResponse{
				Code:		http.StatusBadRequest,
				Message:	r.(string),
			})
			return			
		}
	}() 

	id, _ := uuid.Parse(idParam)
	
	c.JSON(http.StatusOK, h.service.Find(id))
}