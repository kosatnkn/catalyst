package errors

import e "github.com/kosatnkn/catalyst/errors"

// MiddlewareError is the type of errors thrown by middleware.
type MiddlewareError struct {
	*e.BaseError
}

// NewMiddlewareError creates a new MiddlewareError instance.
func NewMiddlewareError(code int, msg string) error {

	return &MiddlewareError{
		BaseError: e.NewBaseError("MiddlewareError", code, msg),
	}
}
