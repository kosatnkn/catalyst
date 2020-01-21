package types

import "fmt"

// MiddlewareError is the type of errors thrown by middleware.
type MiddlewareError struct {
	errType string
	code    int
	msg     string
	details string
}

// NewMiddlewareError creates a new MiddlewareError instance.
func NewMiddlewareError(message string, code int, details string) error {

	return &MiddlewareError{
		errType: "MiddlewareError",
		code:    code,
		msg:     message,
		details: details,
	}
}

// Error returns the MiddlewareError message.
func (e *MiddlewareError) Error() string {
	return fmt.Sprintf("%s|%d|%s|%s", e.errType, e.code, e.msg, e.details)
}
