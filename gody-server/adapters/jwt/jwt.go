package jwt

import (
	"time"

	"github.com/dgrijalva/jwt-go"
)

func GenerateToken(user string) string {
	// Create a new token object, specifying signing method and the claims
	// you would like it to contain.
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"user": user,
		"exp":  time.Now().Add(time.Hour * 2).Unix(),
	})

	// Sign and get the complete encoded token as a string using the secret
	tokenString, _ := token.SignedString([]byte("secret"))

	return tokenString
}

func ValidateToken(validate string) string {
	var tokenCheker string
	tokenString := validate
	claims := jwt.MapClaims{}
	_, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
		return []byte("secret"), nil
	})
	// ... error handling
	if err != nil {
		tokenCheker = "Error"
	} else {
		tokenCheker = "Ok"
	}

	return tokenCheker
}