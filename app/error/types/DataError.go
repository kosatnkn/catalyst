package types

import "fmt"

// DataError is the type of errors thrown by repositories.
type DataError struct {
	msg     string
	code    int
	details string
}

// New creates a new DataError instance.
func (e *DataError) New(message string, code int, details string) error {

	err := &DataError{}

	err.msg = message
	err.code = code
	err.details = details

	return err
}

// Error returns the DataError message.
func (e *DataError) Error() string {
	return fmt.Sprintf("%s|%d|DataError|%s", e.msg, e.code, e.details)
}
