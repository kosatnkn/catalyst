package adapters

import "context"

// DBAdapterInterface is implemented by all database adapters.
type DBAdapterInterface interface {
	// Ping checks wether the database is accessible.
	Ping() error

	// Query runs a query and return the result.
	Query(ctx context.Context, query string, params map[string]interface{}) ([]map[string]interface{}, error)

	// QueryBulk runs a query using an array of parameters and return the combined result.
	//
	// This query is intended to do bulk INSERTS, UPDATES and DELETES.
	// Using this for SELECTS will result in an error.
	QueryBulk(ctx context.Context, query string, params []map[string]interface{}) ([]map[string]interface{}, error)

	// WrapInTx runs the content of the function in a single transaction.
	WrapInTx(ctx context.Context, fn func(ctx context.Context) (interface{}, error)) (interface{}, error)

	// Destruct will close the database adapter releasing all resources.
	Destruct() error
}
