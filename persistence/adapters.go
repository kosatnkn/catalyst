package persistence

import "context"

// DatabaseAdapter is the interface that any database adapter attaching to the service should implement.
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

	// Destruct will close the database adapter releasing all resources.
	Destruct() error

	DatabaseTxAdapter
}

// DatabaseTxAdapter is the interface that any database transaction adapter attaching to the service should implement.
type DatabaseTxAdapter interface {
	// WrapInTx runs the content of the function in a single transaction.
	WrapInTx(ctx context.Context, fn func(ctx context.Context) (any, error)) (any, error)
}

// DatabaseReadinessAdapter is used to report readiness state of the database adapter
// back in to the infrastructure.
type DatabaseReadinessAdapter interface {
	// SetComponent sets the readiness state of the database adapter.
	SetComponent(name string, ready bool)
}
