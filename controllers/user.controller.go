package controllers

import (
	"email/models"
	repo "email/repository"
	"email/utils"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/dgrijalva/jwt-go"
)

type Response models.Response

func HomeController(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello"))
}

func SignInController(w http.ResponseWriter, r *http.Request) {
	var user models.User
	json.NewDecoder(r.Body).Decode(&user)
	w.Header().Set("Content-Type", "application/json")

	res := repo.SignInRepo(&user)
	jsonResponse, _ := json.Marshal(res)

	if !res.Success {
		w.Write(jsonResponse)
		return
	}
	w.Write(jsonResponse)
}

func SignUpController(w http.ResponseWriter, r *http.Request) {
	var user models.User
	json.NewDecoder(r.Body).Decode(&user)
	w.Header().Set("Content-Type", "application/json")

	res := repo.SignUpRepo(&user)
	jsonResponse, _ := json.Marshal(res)

	if !res.Success {
		w.Write(jsonResponse)
		return
	}
	w.Write(jsonResponse)
}

func ComposeController(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	//verify Token
	tokenInfo, err := utils.VerifyAccessToken(r)
	if err != "" {
		jsonResponse, _ := json.Marshal(Response{Message: err, Data: nil, Success: false})
		w.Write(jsonResponse)
		return
	}

	ext := tokenInfo.(jwt.MapClaims)
	userEmail := fmt.Sprintf("%v", ext["email"])

	w.Write([]byte(userEmail))
}

func RefreshTokenController(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	//verify Token
	td, err := utils.VerifyRefreshToken(r)
	if err != "" {
		jsonResponse, _ := json.Marshal(Response{Message: err, Data: nil, Success: false})
		w.Write(jsonResponse)
		return
	}

	tokens := map[string]string{
		"access_token":  td.AccessToken,
		"refresh_token": td.RefreshToken,
	}

	jsonResponse, _ := json.Marshal(Response{Message: "Successfully refresh token.", Data: tokens, Success: true})
	w.Write(jsonResponse)
}
