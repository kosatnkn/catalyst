package errors

import "fmt"

// BaseError is the base error struct that can be used to create different types of errors.
//
// NOTE: Using the BaseError type is not a must. It is given as a continence type that can be used to
//		 derive error structs that can be properly read by the error handler.
type BaseError struct {
	errType string
	code    int
	msg     string
}

// NewBaseError creates a new BaseError.
func NewBaseError(typ string, code int, msg string) *BaseError {

	return &BaseError{
		errType: typ,
		code:    code,
		msg:     msg,
	}
}

// Error returns the error message.
func (e *BaseError) Error() string {
	return fmt.Sprintf("%s|%d|%s", e.errType, e.code, e.msg)
}
