package auth

import (
	"golang.org/x/crypto/bcrypt"
)

var HashCost = 12

func Hash(plain []byte) ([]byte, error) {
	hash, err := bcrypt.GenerateFromPassword(plain, HashCost)
	if err != nil {
		return nil, err
	}

	return hash, nil
}

func VerifyPassword(hash []byte, password []byte) error {
	return bcrypt.CompareHashAndPassword(hash, password)
}
