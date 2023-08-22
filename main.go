package main

import (
	"github.com/Flavio-coutinho/API-Login/routes"
	"github.com/Flavio-coutinho/API-Login/utils"
)

func main() {
	r := routes.SetupRouter()

	r.Use(utils.AuthMiddleware())
	r.Run(":8080")
}