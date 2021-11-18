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
	"github.com/rs/cors"
)

// @title Email services API Documentation.
// @version 1.0.0
// @description A service where users can register and send an email & do live chat.
// @termsOfService http://swagger.io/terms/

// @contact.name Sandeep kumar
// @contact.email sandeepk@gmail.com

// @host localhost:3000
// @BasePath /api/v1

// var addr = flag.String("addr", ":8080", "http service address")

func main() {
	fmt.Println("Startig email services...")

	err := godotenv.Load(".env")
	if err != nil {
		log.Println("Error loading .env file")
	}

	// Database setup
	config.DB = config.SetupDatabase()
	config.DB.AutoMigrate(&models.User{})

	router := routes.SetupRouter()

	// socket .io
	router.Handle("/socket.io/", Socket())
	fs := http.FileServer(http.Dir("static"))
	router.Handle("/", fs)

	// Apply the middleware to the router
	c := cors.New(cors.Options{
		AllowedOrigins:   []string{"*"},
		AllowedMethods:   []string{"GET", "POST"},
		AllowedHeaders:   []string{"Authorization, Content-Type"},
		MaxAge:           50, // in seconds
		AllowCredentials: true,
	})
	handler := c.Handler(router)

	log.Println("Listening on :" + os.Getenv("LOCAL_PORT") + "...")
	error := http.ListenAndServe(":8080", handler)
	if error != nil {
		log.Fatal("Error listening router!")
	}
}
