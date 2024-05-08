package main

import (
	"fmt"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

var SecretKey = []byte("secret")

func generatetoken(user user) (string, error) {

	token := jwt.New(jwt.SigningMethodHS256)
	claims := token.Claims.(jwt.MapClaims)
	claims["username"] = user.name
	claims["password"] = user.password
	claims["exp"] = jwt.NewNumericDate(time.Now().Add(5 * time.Minute))
	claims["admin"] = user.admin

	if user.admin {
		tokenString, err := token.SignedString(SecretKey)
		return tokenString, err
	}

	tokenString, err := token.SignedString(SecretKey)
	fmt.Println(tokenString, err)

	return tokenString, err
}
