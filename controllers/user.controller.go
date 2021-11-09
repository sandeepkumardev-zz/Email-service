package controllers

import (
	"email/models"
	repo "email/repository"
	"email/utils"
	"encoding/json"
	"net/http"

	"github.com/jasonlvhit/gocron"
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

// Email Compose
// @Description Varify user & send an email.
// @Accept  json
// @Produce  json
// @Param email body string true "Send an email"
// @Success 200 {array} Email
// @Router /compose [post]

func EmailComposeController(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	//verify Token
	_, err := utils.VerifyAccessToken(r)
	if err != "" {
		jsonResponse, _ := json.Marshal(Response{Message: err, Data: nil, Success: false})
		w.Write(jsonResponse)
		return
	}

	var T models.EmailTemplate
	json.NewDecoder(r.Body).Decode(&T)

	var resChan = make(chan string)
	var errChan = make(chan string)

	gocron.Every(1).Day().At("17:22").Do(utils.SendEmail, T.To, []byte(T.Message), resChan, errChan)

	select {
	case <-gocron.Start():
	case err := <-errChan:
		jsonResponse, _ := json.Marshal(Response{Message: err, Data: nil, Success: false})
		w.Write(jsonResponse)
	case res := <-resChan:
		jsonResponse, _ := json.Marshal(Response{Message: res, Data: nil, Success: true})
		w.Write(jsonResponse)
	}
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
