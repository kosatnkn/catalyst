package repositories

import e "github.com/kosatnkn/catalyst/v3/internal/errors"

// RepositoryError is the type of errors thrown by repositories.
type RepositoryError struct {
	*e.BaseError
}

// NewRepositoryError creates a new RepositoryError instance.
func NewRepositoryError(code, msg string, cause error) error {
	return &RepositoryError{
		BaseError: e.NewBaseError("RepositoryError", code, msg, cause),
	}
}
