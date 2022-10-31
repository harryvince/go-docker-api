package main

import (
	"os"

	"github.com/gin-gonic/gin"
	"github.com/harryvince/go-docker-api/endpoints"
	"github.com/harryvince/go-docker-api/initializers"
)

func init() {
	initializers.LoadEnvVariables()
}

func main() {
	// Gin setup
	router := gin.Default()
	endpoints.Routes(router)

	// Start server
	PORT := os.Getenv("PORT")
	if PORT == "" {
		panic("PORT is not set in environment variables")
	}
	router.Run(":" + PORT)
}
