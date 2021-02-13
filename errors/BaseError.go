package errors

import "fmt"

// BaseError is the type of errors thrown by business logic.
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

// Error returns the BaseError message.
func (e *BaseError) Error() string {
	return fmt.Sprintf("%s|%d|%s", e.errType, e.code, e.msg)
}
