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

// @title Email services
// @version 1.0
// @description A service where users can register and send an email & do live chat.
// @termsOfService http://swagger.io/terms/
// @contact.name API Support
// @contact.email sandeepypb@gmail.com
// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html
// @host localhost:3000
// @BasePath /

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
