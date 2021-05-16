package errors

import "fmt"

// BaseError is the base error struct that can be used to create different types of errors.
//
// Using the BaseError type is not a must. It is given as a convenience type to
// derive error structs that can be properly read by the error handler.
type BaseError struct {
	typ  string
	code string
	msg  string
	err  error
}

// NewBaseError creates a new BaseError instance.
func NewBaseError(typ, code, msg string, cause error) *BaseError {
	return &BaseError{
		typ:  typ,
		code: code,
		msg:  msg,
		err:  cause,
	}
}

// Error returns the error message.
func (e *BaseError) Error() string {
	return fmt.Sprintf("%s|%s|%s", e.typ, e.code, e.msg)
}

// Unwrap returns the wrapped error.
func (e *BaseError) Unwrap() error {
	return e.err
}
