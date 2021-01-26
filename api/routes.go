package api

import (
	"github.com/gin-gonic/gin"
)

func InitRoutes() *gin.Engine {
	router := gin.Default()

	v1 := router.Group("/api/v1")

	transaction := v1.Group("transaction")
	transaction.GET("/")

	return router
}