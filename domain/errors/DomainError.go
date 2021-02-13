package errors

import "fmt"

// DomainError is the type of errors thrown by business logic.
type DomainError struct {
	errType string
	code    int
	msg     string
	details string
}

// NewDomainError creates a new DomainError.
func NewDomainError(message string, code int, details string) error {

	return &DomainError{
		errType: "DomainError",
		code:    code,
		msg:     message,
		details: details,
	}
}

// Error returns the DomainError message.
func (e *DomainError) Error() string {
	return fmt.Sprintf("%s|%d|%s|%s", e.errType, e.code, e.msg, e.details)
}
