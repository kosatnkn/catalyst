package types

import "fmt"

// MiddlewareError is the type of errors thrown by middleware.
type MiddlewareError struct {
	msg     string
	code    int
	details string
}

// NewMiddlewareError creates a new MiddlewareError instance.
func NewMiddlewareError(message string, code int, details string) error {

	return &MiddlewareError{
		msg:     message,
		code:    code,
		details: details,
	}
}

// Error returns the MiddlewareError message.
func (e *MiddlewareError) Error() string {
	return fmt.Sprintf("%s|%d|MiddlewareError|%s", e.msg, e.code, e.details)
}
