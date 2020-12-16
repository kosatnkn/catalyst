package errors

import "fmt"

// TransformerError is the type of errors thrown by response transformers.
type TransformerError struct {
	errType string
	code    int
	msg     string
	details string
}

// NewTransformerError creates a new TransformerError instance.
func NewTransformerError(message string, code int, details string) error {

	return &TransformerError{
		errType: "TransformerError",
		code:    code,
		msg:     message,
		details: details,
	}
}

// Error returns the TransformerError message.
func (e *TransformerError) Error() string {
	return fmt.Sprintf("%s|%d|%s|%s", e.errType, e.code, e.msg, e.details)
}
