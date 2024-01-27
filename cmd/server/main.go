package main

import (
	"github.com/Dune-Global/unite-backend/internal/config"
	"github.com/Dune-Global/unite-backend/internal/routes/v1"
	"github.com/gin-gonic/gin"
)

func init() {
	config.LoadEnvVariables()
	config.ConnectToDatabase()
}

func main() {
	r := gin.Default()

	routes.ExampleRoutes(r)

	r.Run() 
}