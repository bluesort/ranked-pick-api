package errors

type ForbiddenError struct{}

func (m *ForbiddenError) Error() string {
	return "you do not have access to this resource"
}

func NewForbiddenError() *ForbiddenError {
	return &ForbiddenError{}
}
