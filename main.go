package main

import (
	"github.com/aflores04/financial/transaction/validator"
	"github.com/aflores04/financial/api"
)

func main () {
	apiRouter := api.InitRoutes()

	validator.Register()

	_ = apiRouter.Run("localhost:8000")
}