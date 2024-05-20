package errors

type NotFoundError struct{}

func (m *NotFoundError) Error() string {
	return "not found"
}

func NewNotFoundError() *NotFoundError {
	return &NotFoundError{}
}
