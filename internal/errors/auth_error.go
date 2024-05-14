package errors

type AuthError struct{}

func (m *AuthError) Error() string {
	return "could not authenticate"
}

func NewAuthError() *AuthError {
	return &AuthError{}
}
