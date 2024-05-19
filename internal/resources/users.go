package resources

import (
	"slices"
	"unicode"

	"github.com/carterjackson/ranked-pick-api/internal/errors"
)

type User struct {
	Id          int64  `json:"id"`
	Username    string `json:"username"`
	DisplayName string `json:"display_name,omitempty"`
}

var acceptedPasswordSymbols = []rune{'!', '#', '$', '%', '&', '*', '+', '-', '/', '=', '?', '^', '_', '~', '@'}

func ValidateUsername(username string) error {
	if username == "" {
		return errors.NewInputError("missing username")
	}

	if len(username) > 100 {
		return errors.NewInputError("username must be 100 characters or less")
	}

	// TODO: Validate only letters, numbers, -, and _

	return nil
}

func ValidatePassword(password string) error {
	if password == "" {
		return errors.NewInputError("missing password")
	}

	if len(password) < 8 {
		return errors.NewInputError("password must be at least 8 characters")
	}

	var hasNumber, hasSymbol, hasUppercase bool
	for _, char := range password {
		if unicode.IsSpace(char) {
			return errors.NewInputError("password cannot contain spaces")
		} else if unicode.IsUpper(char) {
			hasUppercase = true
		} else if unicode.IsDigit(char) {
			hasNumber = true
		} else if slices.Contains(acceptedPasswordSymbols, char) {
			hasSymbol = true
		}
	}

	if !hasNumber || !hasSymbol || !hasUppercase {
		return errors.NewInputError("password must contain at least one number, one symbol, and one uppercase letter")
	}

	return nil
}

func ValidateDisplayName(displayName string) error {
	if displayName == "" {
		return errors.NewInputError("missing display name")
	}

	if len(displayName) > 49 {
		return errors.NewInputError("display name must be less than 50 characters")
	}

	return nil
}
