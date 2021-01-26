package main

import (
	"github.com/aflores04/financial/api"
)

func main () {
	apiRouter := api.InitRoutes()

	_ = apiRouter.Run("localhost:8000")
}