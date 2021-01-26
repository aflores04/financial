package main

import (
	"log"
	"os"
)

func main () {
	apiRouter := api.InitRoutes()	

	_ = apiRouter.Run("localhost:8000")
}