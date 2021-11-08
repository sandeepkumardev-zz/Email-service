package main

import (
	"email/config"
	"email/models"
	"email/routes"
	"fmt"
	"log"
	"net/http"
	"os"

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

	router := routes.SetupRouter()
	log.Println("Listening on :" + os.Getenv("LOCAL_PORT") + "...")
	error := http.ListenAndServe(":"+os.Getenv("LOCAL_PORT"), router)
	if error != nil {
		log.Fatal("Error listening router!")
	}
}
