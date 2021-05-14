package services

import (
	"golang.org/x/crypto/bcrypt"
)

func hashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes), err
}

func PasswordIsValid(password, hashedPassword string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
	return err == nil
}

func encryptPassword(password string) string {
	hashed, err := hashPassword(password)
	if err != nil {
		panic(err.Error())
	}
	return hashed
}
