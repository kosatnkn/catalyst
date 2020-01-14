package types

import "fmt"

// RepositoryError is the type of errors thrown by repositories.
type RepositoryError struct {
	msg     string
	code    int
	details string
}

// NewRepositoryError creates a new RepositoryError instance.
func NewRepositoryError(message string, code int, details string) error {

	return &RepositoryError{
		msg:     message,
		code:    code,
		details: details,
	}
}

// Error returns the RepositoryError message.
func (e *RepositoryError) Error() string {
	return fmt.Sprintf("%s|%d|RepositoryError|%s", e.msg, e.code, e.details)
}
