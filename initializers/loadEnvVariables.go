package initializers

import (
	"log"

	"github.com/joho/godotenv"
)

func LoanEnvVariables() {
	err := godotenv.Load()

	if err != nil {
		log.Fatal("Error loading .env file")
	}
}
