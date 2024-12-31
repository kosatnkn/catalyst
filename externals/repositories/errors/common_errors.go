package errors

func ErrQuery(cause error) error {
	return NewRepositoryError("repo-common", "common: error running query", cause)
}
