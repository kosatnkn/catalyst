package types

import "fmt"

// ServiceError is the type of errors thrown by services talking to third party APIs.
type ServiceError struct {
	msg     string
	code    int
	details string
}

// NewServiceError creates a new ServiceError instance.
func NewServiceError(message string, code int, details string) error {

	return &ServiceError{
		msg:     message,
		code:    code,
		details: details,
	}
}

// Error returns the ServiceError message.
func (e *ServiceError) Error() string {
	return fmt.Sprintf("%s|%d|ServiceError|%s", e.msg, e.code, e.details)
}
