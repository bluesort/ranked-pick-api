package auth

import (
	"golang.org/x/crypto/bcrypt"
)

var HashCost = 12

func HashPassword(password []byte) ([]byte, error) {
	hash, err := bcrypt.GenerateFromPassword(password, HashCost)
	if err != nil {
		return nil, err
	}

	return hash, nil
}

func VerifyPassword(hash []byte, password []byte) error {
	return bcrypt.CompareHashAndPassword(hash, password)
}
