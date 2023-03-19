package configs

import (
	"log"

	"github.com/joho/godotenv"
)

func Godotenv() {
	// Landing .env
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
}
