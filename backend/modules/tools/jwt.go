package tools

import (
	"fmt"
	"log"

	jwt "github.com/dgrijalva/jwt-go"
)

var jwtKey = []byte("secret_key_for_ilmli")

func GenerateJWT(input JWTUser) string {
	payload := Payload{
		ID:             input.ID,
		StandardClaims: jwt.StandardClaims{},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, payload)

	tokenString, err := token.SignedString(jwtKey)

	if err != nil {
		log.Fatalf("failed to take string token: %v", err)
	}
	return tokenString
}

func ParseJWT(token string) (JWTUser, error) {

	payload := &Payload{}

	tkn, err := jwt.ParseWithClaims(token, payload, func(token *jwt.Token) (interface{}, error) {
		return jwtKey, nil
	})

	if err != nil {
		fmt.Println(err)
		return JWTUser{ID: 0}, err
	}

	if !tkn.Valid {
		return JWTUser{ID: 0}, err
	}

	return JWTUser{ID: payload.ID}, nil

}
