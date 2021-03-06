package controllers

import (
	Queue "email/hermes-email"
	"email/models"
	repo "email/repository"
	"email/utils"
	"encoding/json"
	"net/http"

	"github.com/jasonlvhit/gocron"
	log "github.com/sirupsen/logrus"
)

type Response models.Response

func HomeController(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello"))
}

// Sign In controller
// @Summary Sign In with credentials.
// @Description A registered user can sign in with their credentials.
// @Tags Sign In
// @Accept  json
// @Produce  json
// @Param user body models.User true "Sign In User"
// @Success 200 {object} models.User
// @Failure 401 {object} object
// @Router /signin [post]
func SignInController(w http.ResponseWriter, r *http.Request) {
	var user models.User
	json.NewDecoder(r.Body).Decode(&user)
	w.Header().Set("Content-Type", "application/json")

	res := repo.SignInRepo(&user)
	jsonResponse, _ := json.Marshal(res)

	if !res.Success {
		log.Warn("Failed to sign in.")
		w.Write(jsonResponse)
		return
	}
	utils.Logger("Signed In as " + user.Email)
	w.Write(jsonResponse)
}

// Sign Up controller
// @Summary Sign Up with credentials.
// @Description A new user can sign up with their email & password.
// @Tags Sign Up
// @Accept  json
// @Produce  json
// @Param user body models.User true "Sign Up User"
// @Success 200 {object} models.User
// @Failure 401 {object} object
// @Router /signup [post]
func SignUpController(w http.ResponseWriter, r *http.Request) {
	var user models.User
	json.NewDecoder(r.Body).Decode(&user)
	w.Header().Set("Content-Type", "application/json")

	res := repo.SignUpRepo(&user)
	jsonResponse, _ := json.Marshal(res)

	if !res.Success {
		log.Warn("Failed to sign up.")
		w.Write(jsonResponse)
		return
	}
	utils.Logger("Signed Up as " + user.Email)
	w.Write(jsonResponse)
}

// Email controller
// @Summary Varify token & send an email.
// @Description You need to signedIn and give a Token in headers then "Send Email" will execute.
// @Tags Email Compose
// @Accept  json
// @Produce  json
// @Param template body models.EmailTemplate true "Send an email"
// @Success 200 {object} models.EmailTemplate
// @Failure 401 {object} object
// @Router /compose [post]
func EmailComposeController(w http.ResponseWriter, r *http.Request) {
	var T models.EmailTemplate
	json.NewDecoder(r.Body).Decode(&T)

	// create the new Queue instance
	Queue.CreateQueue()
	// schedule the CRON job
	s := gocron.NewScheduler()
	s.Every(1).Day().At(T.Time).Do(Queue.Dispatch, T)
	s.Start()
	log.Info("Email in progress ...")

	jsonResponse, _ := json.Marshal(Response{Message: "Email in progress", Data: nil, Success: true})
	w.Write(jsonResponse)
}

// Refresh token controller
// @Summary Varify token & create a new token.
// @Description You need to signedIn and give a Token in headers then "Refresh Token" will execute.
// @Tags Refresh token
// @Accept  json
// @Produce  json
// @Router /refreshToken [post]
func RefreshTokenController(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	//verify Token
	td, err := utils.VerifyRefreshToken(r)
	if err != "" {
		jsonResponse, _ := json.Marshal(Response{Message: err, Data: nil, Success: false})
		log.Error("Error refreshing token")
		w.Write(jsonResponse)
		return
	}

	tokens := map[string]string{
		"access_token":  td.AccessToken,
		"refresh_token": td.RefreshToken,
	}

	jsonResponse, _ := json.Marshal(Response{Message: "Successfully refresh token.", Data: tokens, Success: true})
	utils.Logger("Token refreshed")
	w.Write(jsonResponse)
}
