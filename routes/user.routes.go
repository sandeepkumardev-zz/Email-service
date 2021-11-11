package routes

import (
	Ctrl "email/controllers"
	_ "email/docs"
	M "email/middleware"
	"net/http"
	"os"

	"github.com/gorilla/mux"
	cors "github.com/rs/cors"
	httpSwagger "github.com/swaggo/http-swagger"
)

func SetupRouter() http.Handler {
	router := mux.NewRouter()

	// Api Prefix
	apiPrefix := os.Getenv("API_PREFIX")
	// version 1
	apiVersion := os.Getenv("API_VERSION")

	apiV1 := router.PathPrefix("/" + apiPrefix + "/" + apiVersion).Subrouter()

	getR := apiV1.Methods(http.MethodGet).Subrouter()
	getR.HandleFunc("", Ctrl.HomeController).Schemes("http").Host("localhost:3000")

	postR := apiV1.Methods(http.MethodPost).Subrouter()
	postR.HandleFunc("/signin", Ctrl.SignInController)
	postR.HandleFunc("/signup", Ctrl.SignUpController)
	postR.HandleFunc("/refreshToken", Ctrl.RefreshTokenController)
	postR.HandleFunc("/compose", M.Chain(Ctrl.EmailComposeController, M.VerifyUser()))

	router.PathPrefix("/swagger").Handler(httpSwagger.WrapHandler)

	// logger middleware
	router.Use(M.LoggingMiddleware)

	// Apply the middleware to the router
	c := cors.New(cors.Options{
		AllowedOrigins:   []string{"*"},
		AllowedMethods:   []string{"GET", "POST"},
		AllowedHeaders:   []string{"Authorization, Content-Type"},
		MaxAge:           50, // in seconds
		AllowCredentials: true,
	})
	handler := c.Handler(router)

	return handler
}
