package types

import "fmt"

// ServiceError is the type of errors thrown by services talking to third party APIs.
type ServiceError struct {
	errType string
	code    int
	msg     string
	details string
}

// NewServiceError creates a new ServiceError instance.
func NewServiceError(message string, code int, details string) error {

	return &ServiceError{
		errType: "ServiceError",
		code:    code,
		msg:     message,
		details: details,
	}
}

// Error returns the ServiceError message.
func (e *ServiceError) Error() string {
	return fmt.Sprintf("%s|%d|%s|%s", e.errType, e.code, e.msg, e.details)
}
