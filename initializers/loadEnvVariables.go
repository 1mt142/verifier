package initializers

import (
	"github.com/joho/godotenv"
	"log"
	"os"
)

func LoadEnvVariables() {
	if os.Getenv("ENVIRONMENT") == "docker" {
		return
	}

	err := godotenv.Load()

	if err != nil {
		log.Fatal("Error loading .env file")
	}
}
