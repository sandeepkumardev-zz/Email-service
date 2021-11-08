package controllers

import (
	"email/models"
	repo "email/repository"
	"email/utils"
	"encoding/json"
	"net/http"
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
	tokenInfo, err := utils.VerifyByHeaders(r)
	if err != "" {
		jsonResponse, _ := json.Marshal(Response{Message: err, Data: nil, Success: false})
		w.Write(jsonResponse)
		return
	}
	jsonResponse, _ := json.Marshal(tokenInfo.Claims)
	w.Write([]byte(jsonResponse))
}
