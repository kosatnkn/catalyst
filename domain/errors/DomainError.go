package errors

import e "github.com/kosatnkn/catalyst/errors"

// DomainError is the type of errors thrown by the domain layer.
type DomainError struct {
	*e.BaseError
}

// NewDomainError creates a new DomainError instance.
func NewDomainError(code int, msg string, cause ...error) error {

	return &DomainError{
		BaseError: e.NewBaseError("DomainError", code, msg, cause...),
	}
}
