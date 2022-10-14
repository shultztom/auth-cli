package utils

import (
	"github.com/golang-jwt/jwt"
	"os"
)

var jwtKey = []byte(os.Getenv("JWT_SECRET"))

type Claims struct {
	AppName string `json:"appName"`
	jwt.StandardClaims
}

func CreateToken(appName string) string {
	// Create the JWT claims, which includes the username and expiry time
	claims := &Claims{
		AppName:        appName,
		StandardClaims: jwt.StandardClaims{},
	}
	// Declare the token with the algorithm used for signing, and the claims
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	// Create the JWT string
	tokenString, err := token.SignedString(jwtKey)
	if err != nil {
		// TODO improve error handling
		return "error"
	}
	return tokenString
}
