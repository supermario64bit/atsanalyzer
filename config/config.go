package config

import (
	"log"

	"github.com/joho/godotenv"
)

func LoadEnVFile() {
	err := godotenv.Load(".env.server")
	if err != nil {
		log.Fatal("Error loading .env file. Err: " + err.Error())
	}
}
