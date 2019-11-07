package adapters

import "context"

// DBTxAdapterInterface is implemented by all database transaction adapters.
type DBTxAdapterInterface interface {

	// Wrap runs the content of the function in a single transaction.
	Wrap(ctx context.Context, fn func(ctx context.Context) (interface{}, error)) (interface{}, error)
}
