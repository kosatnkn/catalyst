package types

// ValidationError is the type of errors thrown by the request validator.
type ValidationError struct {
	details string
}

// NewValidationError creates a new ValidationError instance.
func NewValidationError(details string) error {

	err := &ValidationError{}

	err.details = details

	return err
}

// Error returns the ValidationError message.
func (e *ValidationError) Error() string {
	return e.details
}
