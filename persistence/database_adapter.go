package persistence

import "context"

// Context key type to be used with contexts.
type key string

// TxKey is the key used to bind a transaction to context.
const DatabaseTxKey key = "tx"

const (
	DatabaseAffectedRows string = "affected_rows"
	DatabaseLastInsertID string = "last_insert_id"
)

// DatabaseAdapter is implemented by all database adapters.
type DatabaseAdapter interface {
	// Ping checks wether the database is accessible.
	Ping() error

	// Query runs a query and return the result.
	Query(ctx context.Context, query string, params map[string]any) ([]map[string]any, error)

	// QueryBulk runs a query using an array of parameters and return the combined result.
	//
	// This query is intended to do bulk INSERTS, UPDATES and DELETES.
	// Using this for SELECTS will result in an error.
	QueryBulk(ctx context.Context, query string, params []map[string]any) ([]map[string]any, error)

	// WrapInTx runs the content of the function in a single transaction.
	WrapInTx(ctx context.Context, fn func(ctx context.Context) (any, error)) (any, error)

	// Destruct will close the database adapter releasing all resources.
	Destruct() error
}
