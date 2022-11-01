package main

import (
	"os"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/harryvince/go-docker-api/Backend/endpoints"
	"github.com/harryvince/go-docker-api/Backend/initializers"
)

func init() {
	initializers.LoadEnvVariables()
	initializers.PullAllImages()
}

func main() {
	// Gin setup
	router := gin.Default()
	router.Use(cors.New(cors.Config{
		AllowOrigins:  []string{"*"},
		AllowHeaders:  []string{"Origin"},
		ExposeHeaders: []string{"Content-Length"},
	}))
	endpoints.Routes(router)

	// Start server
	PORT := os.Getenv("PORT")
	if PORT == "" {
		panic("PORT is not set in environment variables")
	}
	router.Run(":" + PORT)
}
