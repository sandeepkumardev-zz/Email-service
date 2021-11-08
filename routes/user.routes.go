package routes

import (
	"github.com/gorilla/mux"

	Ctrl "email/controllers"
)

func SetupRouter() *mux.Router {
	router := mux.NewRouter()

	router.HandleFunc("/", Ctrl.HomeController).Methods("GET")

	return router
}
