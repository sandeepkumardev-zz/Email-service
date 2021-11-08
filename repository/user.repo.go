package repository

import (
	"email/config"
	"email/models"
)

type Response models.Response

func SignInRepo(user *models.User) Response {
	var pswd = user.Password
	if err := config.DB.Where("email = ?", user.Email).First(user).Error; err != nil {
		return Response{Message: "SignedIn failed!", Data: nil, Success: false}
	}
	if user.Password != pswd {
		return Response{Message: "Wrong password!", Data: nil, Success: false}
	}
	return Response{Message: "SignedIn successfully", Data: nil, Success: true}
}

func SignUpRepo(user models.User) Response {
	if err := config.DB.Create(user).Error; err != nil {
		return Response{Message: "SignedUp failed!", Data: nil, Success: false}
	}
	return Response{Message: "SignedUp successfully", Data: nil, Success: true}
}
