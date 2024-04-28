package errors

type InputError struct {
	Message string
}

func (m *InputError) Error() string {
	return m.Message
}

func NewInputError(message string) *InputError {
	return &InputError{Message: message}
}
