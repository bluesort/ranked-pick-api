package auth

import (
	"bytes"
	"crypto/rand"
	"encoding/base64"
	"errors"

	"golang.org/x/crypto/argon2"
)

// var HashCost = 12

const (
	HashTime    = 1
	HashMemory  = 64 * 1024
	HashThreads = 4
	HashKeyLen  = 32
	HashSaltLen = 32
)

func Hash(plain []byte) (string, error) {
	salt := make([]byte, HashSaltLen)
	_, err := rand.Read(salt)
	if err != nil {
		return "", err
	}

	key := argon2.IDKey(plain, salt, HashTime, HashMemory, HashThreads, HashKeyLen)
	toHash := append(key, salt...)
	b64 := base64.StdEncoding.EncodeToString(toHash)

	return b64, nil
}

func VerifyPassword(encodedHash string, password string) error {
	decoded, err := base64.StdEncoding.DecodeString(string(encodedHash))
	if err != nil {
		return err
	}
	storedHash := decoded[:HashKeyLen]
	salt := decoded[HashKeyLen:]

	passwordHash := argon2.IDKey([]byte(password), salt, HashTime, HashMemory, HashThreads, HashKeyLen)
	if !bytes.Equal(storedHash, passwordHash) {
		// TODO: auth error
		return errors.New("invalid password")
	}

	return nil
}
