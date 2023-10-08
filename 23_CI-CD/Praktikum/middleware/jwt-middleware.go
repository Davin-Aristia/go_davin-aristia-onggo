package middleware

import (
	"cicd/constant"
	"time"

	"github.com/golang-jwt/jwt"
)

func CreateToken(userId int) (string, error) {
	claims := jwt.MapClaims{}
	// token kedua (payload)
	claims["authorized"] = true
	claims["userId"] = userId
	claims["exp"] = time.Now().Add(time.Hour * 1).Unix() //Token expires after 1 hour
	// token pertama (header)
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	// return bersama token ketiga (dengan secret key)
	return token.SignedString([]byte(constant.SECRET_JWT))
}
