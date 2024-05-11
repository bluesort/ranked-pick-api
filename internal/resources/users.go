package resources

import (
	"net/mail"
	"slices"
	"unicode"

	"github.com/carterjackson/ranked-pick-api/internal/errors"
)

type User struct {
	Id          int64  `json:"id"`
	Email       string `json:"email"`
	DisplayName string `json:"display_name,omitempty"`
}

var acceptedPasswordSymbols = []rune{'!', '#', '$', '%', '&', '*', '+', '-', '/', '=', '?', '^', '_', '~', '@'}

func ValidateEmail(email string) error {
	if email == "" {
		return errors.NewInputError("missing email")
	}

	if len(email) > 299 {
		return errors.NewInputError("email must be less than 300 characters")
	}

	_, err := mail.ParseAddress(email)
	if err != nil {
		return errors.NewInputError("invalid email")
	}

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
