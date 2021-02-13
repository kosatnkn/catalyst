package errors

// RepositoryError is the type of errors thrown by repositories.
import e "github.com/kosatnkn/catalyst/errors"

// RepositoryError is the type of errors thrown by middleware.
type RepositoryError struct {
	*e.BaseError
}

// NewRepositoryError creates a new RepositoryError instance.
func NewRepositoryError(code int, msg string) error {

	return &RepositoryError{
		BaseError: e.NewBaseError("RepositoryError", code, msg),
	}
}
