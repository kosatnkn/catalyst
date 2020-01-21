package types

import "fmt"

// ServerError is the type of errors thrown by the framework while booting.
type ServerError struct {
	errType string
	code    int
	msg     string
	details string
}

// NewServerError creates a new ServerError instance.
func NewServerError(message string, code int, details string) error {

	return &ServerError{
		errType: "ServerError",
		code:    code,
		msg:     message,
		details: details,
	}
}

// Error returns the ServerError message.
func (e *ServerError) Error() string {
	return fmt.Sprintf("%s|%d|%s|%s", e.errType, e.code, e.msg, e.details)
}
