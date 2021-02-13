package errors

import "fmt"

// RepositoryError is the type of errors thrown by repositories.
type RepositoryError struct {
	errType string
	code    int
	msg     string
}

// NewRepositoryError creates a new RepositoryError instance.
func NewRepositoryError(code int, msg string) error {

	return &RepositoryError{
		errType: "RepositoryError",
		code:    code,
		msg:     msg,
	}
}

// Error returns the RepositoryError message.
func (e *RepositoryError) Error() string {
	return fmt.Sprintf("%s|%d|%s", e.errType, e.code, e.msg)
}
