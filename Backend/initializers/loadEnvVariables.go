package initializers

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

// LoadEnvVariables loads environment variables from .env file
func LoadEnvVariables() {
	err := os.Getenv("ENVIRONMENT")
	if err != "PRODUCTION" {
		err := godotenv.Load()
		if err != nil {
			log.Fatal("Error loading .env file")
		} else {
			log.Println("Loaded .env file")
		}
	}
}
