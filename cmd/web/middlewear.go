package main

import (
	"log"
	"net/http"
	"strconv"

	"github.com/golang-jwt/jwt"
)

func jwtValidate(tokenString string) (jwt.MapClaims, error) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, ErrUnexpectedSigningMethod
		}

		return SecretKey, nil
	})

	if err != nil || !token.Valid {
		return nil, err
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		return nil, err
	}

	return claims, nil
}

func userAuth(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		tokenString := r.Header.Get("Authorization")
		if tokenString == "" {
			http.Error(w, ErrAuthHeaderRequired.Error(), http.StatusUnauthorized)
			return
		}

		claims, err := jwtValidate(tokenString)
		if err != nil {
			log.Print("Jwt validation err", err)
			http.Error(w, ErrInvalidOrExpiredToken.Error(), http.StatusUnauthorized)
			return
		}

		setHeaders(r, claims)

		next.ServeHTTP(w, r)
	}
}

func adminAuth(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		tokenString := r.Header.Get("Authorization")
		if tokenString == "" {
			http.Error(w, ErrAuthHeaderRequired.Error(), http.StatusUnauthorized)
			return
		}

		claims, err := jwtValidate(tokenString)
		if err != nil {
			http.Error(w, ErrInvalidOrExpiredToken.Error(), http.StatusUnauthorized)
			return
		}

		if claims["admin"] == false {
			http.Error(w, ErrUnAuthorizedUser.Error(), http.StatusUnauthorized)
			return
		}

		setHeaders(r, claims)
		next.ServeHTTP(w, r)
	}
}

func setHeaders(r *http.Request, claims jwt.MapClaims) {
	r.Header.Set("username", claims["username"].(string))
	r.Header.Set("password", claims["password"].(string))
	r.Header.Set("admin", strconv.FormatBool(claims["admin"].(bool)))
}
