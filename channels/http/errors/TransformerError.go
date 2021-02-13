package errors

import "fmt"

// TransformerError is the type of errors thrown by response transformers.
type TransformerError struct {
	errType string
	code    int
	msg     string
}

// NewTransformerError creates a new TransformerError instance.
func NewTransformerError(code int, msg string) error {

	return &TransformerError{
		errType: "TransformerError",
		code:    code,
		msg:     msg,
	}
}

// Error returns the TransformerError message.
func (e *TransformerError) Error() string {
	return fmt.Sprintf("%s|%d|%s", e.errType, e.code, e.msg)
}
