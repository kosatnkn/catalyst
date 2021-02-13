package errors

import "fmt"

// MiddlewareError is the type of errors thrown by middleware.
type MiddlewareError struct {
	errType string
	code    int
	msg     string
}

// NewMiddlewareError creates a new MiddlewareError instance.
func NewMiddlewareError(code int, msg string) error {

	return &MiddlewareError{
		errType: "MiddlewareError",
		code:    code,
		msg:     msg,
	}
}

// Error returns the MiddlewareError message.
func (e *MiddlewareError) Error() string {
	return fmt.Sprintf("%s|%d|%s", e.errType, e.code, e.msg)
}
