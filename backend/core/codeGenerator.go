package core

import (
	"math/rand"
)

var alphabet = "ABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"

// GenerateCode generates a six digits random alphanumeric code
func GenerateCode() string {
	b := make([]byte, 6)
	for i := range b {
		b[i] = alphabet[rand.Intn(len(alphabet))]
	}
	return string(b)
}
