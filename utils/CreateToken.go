package utils

import (
	"os"
	"time"

	"github.com/dgrijalva/jwt-go"
)

func CreateToken(email string, expireTokenTime time.Time) (string, error) {
	atClaims := jwt.MapClaims{}
	atClaims["expiresAt"] = expireTokenTime.Unix()
	atClaims["email"] = email
	atClaims["authorized"] = true

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, atClaims)
	tokenSting, err := token.SignedString([]byte(os.Getenv("JWT_SECRET")))
	if err != nil {
		return "", err

	}
	return tokenSting, nil
}
