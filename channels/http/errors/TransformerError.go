package errors

// TransformerError is the type of errors thrown by response transformers.
import e "github.com/kosatnkn/catalyst/errors"

// TransformerError is the type of errors thrown by middleware.
type TransformerError struct {
	*e.BaseError
}

// NewTransformerError creates a new TransformerError instance.
func NewTransformerError(code int, msg string) error {

	return &TransformerError{
		BaseError: e.NewBaseError("TransformerError", code, msg),
	}
}
