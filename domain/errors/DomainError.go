package errors

import e "github.com/kosatnkn/catalyst/errors"

// DomainError is the type of errors thrown by business logic.
type DomainError struct {
	*e.BaseError
}

// NewDomainError creates a new DomainError instance.
func NewDomainError(code int, msg string) error {

	return &DomainError{
		BaseError: e.NewBaseError("DomainError", code, msg),
	}
}
