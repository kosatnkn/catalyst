package errors

import e "github.com/kosatnkn/catalyst/errors"

// ServiceError is the type of errors thrown by services talking to third party APIs.
type ServiceError struct {
	*e.BaseError
}

// NewServiceError creates a new ServiceError instance.
func NewServiceError(code int, msg string, cause ...error) error {

	return &ServiceError{
		BaseError: e.NewBaseError("ServiceError", code, msg, cause...),
	}
}
