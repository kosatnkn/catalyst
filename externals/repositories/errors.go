package repositories

func ErrQuery(cause error) error {
	return NewRepositoryError("repo-common", "repo: error running query", cause)
}
