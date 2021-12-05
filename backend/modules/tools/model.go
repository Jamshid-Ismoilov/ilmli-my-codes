package tools

import "github.com/dgrijalva/jwt-go"

type JWTUser struct {
	ID int
}

type Payload struct {
	ID int
	jwt.StandardClaims
}
