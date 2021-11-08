package main

import (
	"email/config"
	"email/models"
	"fmt"
	"log"

	"github.com/joho/godotenv"
)

func main() {
	fmt.Println("Startig email services...")

	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	config.DB = config.SetupDatabase()
	config.DB.AutoMigrate(&models.User{})
}
