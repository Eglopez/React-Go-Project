package services

import (
	"errors"
	"time"

	"github.com/LKezHn/React-Go-Project/core/models"
	"github.com/dgrijalva/jwt-go"
)

func NewToken(id string) string {

	expirationTime := time.Now().Add(3 * time.Minute)

	claims := models.JWT{
		Id: id,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expirationTime.Unix(),
			Issuer:    "myBackend",
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	signedToken, _ := token.SignedString([]byte("mysupersecret"))

	return signedToken
}

func ValidateToken(jwtFromHeader string) (string, error) {
	token, _ := jwt.ParseWithClaims(
		jwtFromHeader,
		&models.JWT{},
		func(token *jwt.Token) (interface{}, error) {
			return []byte("mysupersecret"), nil
		},
	)

	claims, ok := token.Claims.(*models.JWT)

	if !ok {
		return "0", errors.New("Invalid Token")
	}

	if time.Now().After(time.Unix(int64(claims.ExpiresAt), 0)) {
		return "0", errors.New("Token has expired")
	}

	id := claims.Id

	return id, nil

}
