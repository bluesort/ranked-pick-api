package auth

import (
	"bytes"
	"crypto/rand"
	"encoding/base64"

	"github.com/carterjackson/ranked-pick-api/internal/errors"
	"golang.org/x/crypto/argon2"
)

const (
	HashTime    = 1
	HashMemory  = 64 * 1024
	HashThreads = 4
	HashKeyLen  = 32
	HashSaltLen = 32
)

func Hash(plain string) (string, error) {
	salt := make([]byte, HashSaltLen)
	_, err := rand.Read(salt)
	if err != nil {
		return "", err
	}

	key := argon2.IDKey([]byte(plain), salt, HashTime, HashMemory, HashThreads, HashKeyLen)
	toHash := append(key, salt...)
	b64 := base64.StdEncoding.EncodeToString(toHash)

	return b64, nil
}

func VerifyPlainWithHash(plain string, encodedHash string) error {
	decoded, err := base64.StdEncoding.DecodeString(string(encodedHash))
	if err != nil {
		return err
	}
	storedHash := decoded[:HashKeyLen]
	salt := decoded[HashKeyLen:]

	plainHash := argon2.IDKey([]byte(plain), salt, HashTime, HashMemory, HashThreads, HashKeyLen)
	if !bytes.Equal(storedHash, plainHash) {
		return errors.NewAuthError()
	}

	return nil
}
