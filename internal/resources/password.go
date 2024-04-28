package resources

import (
	"golang.org/x/crypto/bcrypt"
)

var PasswordHashCost = 12

func HashPassword(password []byte) ([]byte, error) {
	hash, err := bcrypt.GenerateFromPassword(password, PasswordHashCost)
	if err != nil {
		return nil, err
	}

	return hash, nil
}

func VerifyPassword(hash []byte, password []byte) error {
	return bcrypt.CompareHashAndPassword(hash, password)
}
