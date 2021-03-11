package errors

import "fmt"

// BaseError is the base error struct that can be used to create different types of errors.
//
// NOTE:
// Using the BaseError type is not a must. It is given as a convenience type to
// derive error structs that can be properly read by the error handler.
type BaseError struct {
	errType string
	code    string
	msg     string
	err     error
}

// NewBaseError creates a new BaseError instance.
func NewBaseError(typ, code, msg string, cause ...error) *BaseError {

	e := &BaseError{
		errType: typ,
		code:    code,
		msg:     msg,
	}

	if len(cause) > 0 {
		e.err = cause[0]
	}

	return e
}

// Error returns the error message.
func (e *BaseError) Error() string {
	return fmt.Sprintf("%s|%s|%s", e.errType, e.code, e.msg)
}

// Unwrap returns the wrapped error.
func (e *BaseError) Unwrap() error {
	return e.err
}
