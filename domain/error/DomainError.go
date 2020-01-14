package error

import "fmt"

// DomainError is the type of errors thrown by business logic.
type DomainError struct {
	msg     string
	code    int
	details string
}

// NewDomainError creates a new DomainError.
func NewDomainError(message string, code int, details string) error {

	return &DomainError{
		msg:     message,
		code:    code,
		details: details,
	}
}

// Error returns the DomainError message.
func (e *DomainError) Error() string {
	return fmt.Sprintf("%s|%d|DomainError|%s", e.msg, e.code, e.details)
}
