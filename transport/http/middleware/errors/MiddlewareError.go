package errors

import e "github.com/kosatnkn/catalyst/v3/internal/errors"

// MiddlewareError is the type of errors thrown by middleware.
type MiddlewareError struct {
	*e.BaseError
}

// NewMiddlewareError creates a new MiddlewareError instance.
func NewMiddlewareError(code, msg string, errs error) error {
	return &MiddlewareError{
		BaseError: e.NewBaseError("MiddlewareError", code, msg, errs),
	}
}
