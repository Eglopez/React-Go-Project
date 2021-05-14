package models

import "github.com/dgrijalva/jwt-go"

type JWT struct {
	Id string `json:"id"`
	jwt.StandardClaims
}
