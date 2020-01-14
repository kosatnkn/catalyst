package types

import "fmt"

// ServerError is the type of errors thrown by the framework while booting.
type ServerError struct {
	msg     string
	code    int
	details string
}

// NewServerError creates a new ServerError instance.
func NewServerError(message string, code int, details string) error {

	return &ServerError{
		msg:     message,
		code:    code,
		details: details,
	}
}

// Error returns the ServerError message.
func (e *ServerError) Error() string {
	return fmt.Sprintf("%s|%d|ServerError|%s", e.msg, e.code, e.details)
}
