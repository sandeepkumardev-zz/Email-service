package utils

import (
	M "email/models"
	"fmt"
	"net/http"
	"os"
	"time"

	"github.com/dgrijalva/jwt-go"
	// str2duration "github.com/xhit/go-str2duration/v2"
)

var td = &M.TokenDetails{}

func CreateToken(email string) (*M.TokenDetails, error) {
	// min, _ := str2duration.ParseDuration(os.Getenv("EXPIRE_ACCESS_SECRET"))
	td.AtExpires = time.Now().Add(time.Minute * 1).Unix()
	td.RtExpires = time.Now().Add(time.Minute * 2).Unix()

	var err error
	//Creating Access Token
	atClaims := jwt.MapClaims{}
	atClaims["expiresAt"] = td.AtExpires
	atClaims["email"] = email
	atClaims["authorized"] = true
	at := jwt.NewWithClaims(jwt.SigningMethodHS256, atClaims)
	td.AccessToken, err = at.SignedString([]byte(os.Getenv("ACCESS_SECRET")))
	if err != nil {
		return nil, err
	}

	//Creating Refresh Token
	rtClaims := jwt.MapClaims{}
	rtClaims["expiresAt"] = td.RtExpires
	rtClaims["email"] = email
	rtClaims["authorized"] = true
	rt := jwt.NewWithClaims(jwt.SigningMethodHS256, rtClaims)
	td.RefreshToken, err = rt.SignedString([]byte(os.Getenv("REFRESH_SECRET")))
	if err != nil {
		return nil, err
	}

	return td, nil
}

func VerifyAccessToken(req *http.Request) (jwt.Claims, string) {
	if data := req.Header.Get("Authorization"); data != "" {
		token, err := verifyToken(data, []byte(os.Getenv("ACCESS_SECRET")))
		if err != "" {
			return nil, err
		}

		// extract expiresTime from token
		ext := token.Claims.(jwt.MapClaims)
		expiresTime := ext["expiresAt"]

		if int64(expiresTime.(float64)) < time.Now().Unix() {
			return nil, "Token expired!"
		}
		return token.Claims, ""
	} else {
		return nil, "You are not logged In"
	}
}

func VerifyRefreshToken(req *http.Request) (*M.TokenDetails, string) {
	if data := req.Header.Get("Authorization"); data != "" {
		token, err := verifyToken(data, []byte(os.Getenv("REFRESH_SECRET")))
		if err != "" {
			return nil, err
		}

		// extract email from token
		ext := token.Claims.(jwt.MapClaims)
		email := fmt.Sprintf("%v", ext["email"])

		// create new tokens
		td, Terr := CreateToken(email)
		if Terr != nil {
			return nil, "Something went wrong!"
		}
		return td, ""
	} else {
		return nil, "You are not logged In"
	}
}

func verifyToken(data string, secret []byte) (*jwt.Token, string) {
	token, err := jwt.Parse(data, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return secret, nil
	})

	if err != nil {
		return nil, "Invalid Token!"
	}

	return token, ""
}
