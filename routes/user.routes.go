package routes

import (
	"github.com/gorilla/mux"

	Ctrl "email/controllers"

	_ "email/docs"

	httpSwagger "github.com/swaggo/http-swagger"
)

// @title Email services API
// @version 1.0
// @description A service where users can register and send an email & do live chat.
// @termsOfService http://swagger.io/terms/
// @contact.name API Support
// @contact.email sandeepk@gmail.com
// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html
// @host localhost:3000
// @BasePath /

func SetupRouter() *mux.Router {
	router := mux.NewRouter()

	router.HandleFunc("/", Ctrl.HomeController).Methods("GET")

	router.HandleFunc("/signin", Ctrl.SignInController).Methods("Post")
	router.HandleFunc("/signup", Ctrl.SignUpController).Methods("Post")

	router.HandleFunc("/compose", Ctrl.EmailComposeController).Methods("Post")
	router.HandleFunc("/refreshToken", Ctrl.RefreshTokenController).Methods("Post")

	router.PathPrefix("/swagger").Handler(httpSwagger.WrapHandler)
	return router
}
