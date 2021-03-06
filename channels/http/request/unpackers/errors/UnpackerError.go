package errors

import e "github.com/kosatnkn/catalyst/v2/errors"

// UnpackerError is the type of errors thrown by response transformers.
type UnpackerError struct {
	*e.BaseError
}

// NewUnpackerError creates a new UnpackerError instance.
func NewUnpackerError(code string, msg string, cause ...error) error {

	return &UnpackerError{
		BaseError: e.NewBaseError("UnpackerError", code, msg, cause...),
	}
}
