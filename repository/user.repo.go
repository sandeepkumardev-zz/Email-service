package repository

import (
	"email/config"
	"email/models"
	"email/utils"
	"time"

	"golang.org/x/crypto/bcrypt"
)

type Response models.Response
type Token models.Token

var err error

func SignInRepo(user *models.User) Response {
	var pswd = user.Password
	if err = config.DB.Where("email = ?", user.Email).First(user).Error; err != nil {
		return Response{Message: "Record not found!", Data: nil, Success: false}
	}
	if err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(pswd)); err != nil {
		return Response{Message: "Wrong password!", Data: nil, Success: false}
	}

	// genrate JWT token
	expireTokenTime := time.Now().Add(time.Minute * 10)
	tokenSting, err := utils.CreateToken(user.Email, expireTokenTime)
	if err != nil {
		return Response{Message: "Something went wrong!", Data: nil, Success: false}
	}

	return Response{Message: "SignedIn successfully", Data: Token{Token: tokenSting}, Success: true}
}

func SignUpRepo(user *models.User) Response {
	if err = config.DB.Where("email = ?", user.Email).First(user).Error; err != nil {
		// hash password
		hash, hashErr := bcrypt.GenerateFromPassword([]byte(user.Password), 12)
		if hashErr != nil {
			return Response{Message: "Something went wrong!", Data: nil, Success: false}
		}

		user.Password = string(hash)

		if err = config.DB.Create(user).Error; err != nil {
			return Response{Message: "SignedUp failed!", Data: nil, Success: false}
		}
		return Response{Message: "SignedUp successfully", Data: nil, Success: true}
	}
	return Response{Message: "Record already exists!", Data: nil, Success: false}
}
