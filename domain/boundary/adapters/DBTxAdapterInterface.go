package adapters

import "database/sql"

// DBTxAdapterInterface is implemented by all database transaction adapters.
type DBTxAdapterInterface interface {

	// Wrap runs the content of the function in a single transaction.
	Wrap(fn func(*sql.Tx) (interface{}, error)) (interface{}, error)
}
