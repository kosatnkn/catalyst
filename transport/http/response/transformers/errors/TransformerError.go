package errors

import e "github.com/kosatnkn/catalyst/v3/errors"

// TransformerError is the type of errors thrown by response transformers.
type TransformerError struct {
	*e.BaseError
}

// NewTransformerError creates a new TransformerError instance.
func NewTransformerError(code, msg string, cause error) error {
	return &TransformerError{
		BaseError: e.NewBaseError("TransformerError", code, msg, cause),
	}
}
