package errors

import "fmt"

// DomainError is the type of errors thrown by business logic.
type DomainError struct {
	errType string
	code    int
	msg     string
}

// NewDomainError creates a new DomainError.
func NewDomainError(code int, msg string) error {

	return &DomainError{
		errType: "DomainError",
		code:    code,
		msg:     msg,
	}
}

// Error returns the DomainError message.
func (e *DomainError) Error() string {
	return fmt.Sprintf("%s|%d|%s", e.errType, e.code, e.msg)
}
