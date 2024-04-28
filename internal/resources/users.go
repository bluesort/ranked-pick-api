package resources

import "errors"

func ValidateEmail(email string) error {
	if email == "" {
		return errors.New("missing email")
	}
	return nil
}

func ValidatePassword(password string) error {
	if password == "" {
		return errors.New("missing password")
	}
	return nil
}

func ValidateDisplayName(displayName string) error {
	if displayName == "" {
		return errors.New("missing display name")
	}
	return nil
}
