package middleware

import (
	"email/models"
	"email/utils"
	"encoding/json"
	"net/http"
)

type Middleware func(http.HandlerFunc) http.HandlerFunc

func VerifyUser() Middleware {
	// Create a new Middleware
	return func(next http.HandlerFunc) http.HandlerFunc {
		// Define the http.HandlerFunc
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			w.Header().Set("Content-Type", "application/json")
			//verify Token
			_, err := utils.VerifyAccessToken(r)
			if err != "" {
				jsonResponse, _ := json.Marshal(models.Response{Message: err, Data: nil, Success: false})
				w.Write(jsonResponse)
				return
			}
			next.ServeHTTP(w, r)
		})
	}
}

// Chain applies middlewares to a http.HandlerFunc
func Chain(f http.HandlerFunc, middlewares ...Middleware) http.HandlerFunc {
	for _, m := range middlewares {
		f = m(f)
	}
	return f
}
