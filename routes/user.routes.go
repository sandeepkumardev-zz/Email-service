package routes

import (
	"github.com/gorilla/mux"

	Ctrl "email/controllers"
)

func SetupRouter() *mux.Router {
	router := mux.NewRouter()

	router.HandleFunc("/", Ctrl.HomeController).Methods("GET")

	router.HandleFunc("/signin", Ctrl.SignInController).Methods("Post")
	router.HandleFunc("/signup", Ctrl.SignUpController).Methods("Post")

	router.HandleFunc("/compose", Ctrl.ComposeController).Methods("Post")
	router.HandleFunc("/refreshToken", Ctrl.RefreshTokenController).Methods("Post")

	return router
}
