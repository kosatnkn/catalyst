package errors

import "fmt"

// ServiceError is the type of errors thrown by services talking to third party APIs.
type ServiceError struct {
	errType string
	code    int
	msg     string
}

// NewServiceError creates a new ServiceError instance.
func NewServiceError(code int, msg string) error {

	return &ServiceError{
		errType: "ServiceError",
		code:    code,
		msg:     msg,
	}
}

// Error returns the ServiceError message.
func (e *ServiceError) Error() string {
	return fmt.Sprintf("%s|%d|%s", e.errType, e.code, e.msg)
}
