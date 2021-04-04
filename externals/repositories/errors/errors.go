package errors

func ErrQuery(cause error) error {
	return NewRepositoryError("100", "Error running query", cause)
}
