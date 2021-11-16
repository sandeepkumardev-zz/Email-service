package routes

import (
	Ctrl "email/controllers"
	_ "email/docs"
	M "email/middleware"
	"net/http"
	"os"

	"github.com/gorilla/mux"
	httpSwagger "github.com/swaggo/http-swagger"
)

func SetupRouter() *mux.Router {
	router := mux.NewRouter()

	// Api Prefix
	apiPrefix := os.Getenv("API_PREFIX")
	// version 1
	apiVersion := os.Getenv("API_VERSION")

	apiV1 := router.PathPrefix("/" + apiPrefix + "/" + apiVersion).Subrouter()

	// getR := apiV1.Methods(http.MethodGet).Subrouter()
	// getR.HandleFunc("", Ctrl.HomeController).Schemes("http").Host("localhost:3000")

	postR := apiV1.Methods(http.MethodPost).Subrouter()
	postR.HandleFunc("/signin", Ctrl.SignInController)
	postR.HandleFunc("/signup", Ctrl.SignUpController)
	postR.HandleFunc("/refreshToken", Ctrl.RefreshTokenController)
	postR.HandleFunc("/compose", M.Chain(Ctrl.EmailComposeController, M.VerifyUser()))

	router.PathPrefix("/swagger").Handler(httpSwagger.WrapHandler)

	// logger middleware
	router.Use(M.LoggingMiddleware)

	return router
}
