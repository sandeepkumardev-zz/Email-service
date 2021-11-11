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

// @title Email services API Documentation.
// @version 1.0.0
// @description A service where users can register and send an email & do live chat.
// @termsOfService http://swagger.io/terms/

// @contact.name Sandeep kumar
// @contact.email sandeepk@gmail.com

// @host localhost:3000
// @BasePath /api/v1

func main() {
	fmt.Println("Startig email services...")

	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	config.DB = config.SetupDatabase()
	config.DB.AutoMigrate(&models.User{})

	handler := routes.SetupRouter()
	log.Println("Listening on :" + os.Getenv("LOCAL_PORT") + "...")
	error := http.ListenAndServe(":"+os.Getenv("LOCAL_PORT"), handler)
	if error != nil {
		log.Fatal("Error listening router!")
	}
}
