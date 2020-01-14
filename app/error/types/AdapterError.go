package types

import "fmt"

// AdapterError is the type of errors thrown by adapters.
type AdapterError struct {
	msg     string
	code    int
	details string
}

// NewAdapterError creates a new AdapterError instance.
func NewAdapterError(message string, code int, details string) error {

	return &AdapterError{
		msg:     message,
		code:    code,
		details: details,
	}
}

// Error returns the AdapterError message.
func (e *AdapterError) Error() string {
	return fmt.Sprintf("%s|%d|AdapterError|%s", e.msg, e.code, e.details)
}
