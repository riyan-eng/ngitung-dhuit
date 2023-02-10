package config

import (
	"fmt"
	"log"

	"github.com/joho/godotenv"
)

func Environment() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	fmt.Println("Success load environment")
}
