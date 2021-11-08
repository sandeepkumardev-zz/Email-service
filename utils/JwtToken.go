package utils

import (
	"fmt"
	"net/http"
	"os"
	"time"

	"github.com/dgrijalva/jwt-go"
	str2duration "github.com/xhit/go-str2duration/v2"
)

var AtClaims = jwt.MapClaims{}

func CreateToken(email string) (string, error) {
	min, _ := str2duration.ParseDuration(os.Getenv("EXPIRE_SECRET"))
	expireTokenTime := time.Now().Add(time.Minute * min)
	AtClaims["expiresAt"] = expireTokenTime.Unix()
	AtClaims["email"] = email
	AtClaims["authorized"] = true

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, AtClaims)
	tokenSting, err := token.SignedString([]byte(os.Getenv("JWT_SECRET")))
	if err != nil {
		return "", err

	}
	return tokenSting, nil
}

func VerifyByHeaders(req *http.Request) (*jwt.Token, string) {
	if data := req.Header.Get("Authorization"); data != "" {
		token, err := verifyToken(data)
		if err != "" {
			return nil, err
		}
		return token, ""
	} else {
		return nil, "You are not logged In"
	}
}

func verifyToken(data string) (*jwt.Token, string) {
	token, err := jwt.ParseWithClaims(data, &AtClaims, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(os.Getenv("JWT_SECRET")), nil
	})
	if err != nil {
		return nil, "Invalid Token!"
	}

	// v, _ := err.(*jwt.ValidationError)

	// if v.Errors == jwt.ValidationErrorExpired && int64(AtClaims["expiresAt"].(float64)) > time.Now().Unix()-(86400*14) {
	// 	return nil, "Token expired"
	// }

	return token, ""
}
