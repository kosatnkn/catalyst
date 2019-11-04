package adapters

import "database/sql"

// DBAdapterInterface is implemented by all database adapters.
type DBAdapterInterface interface {

	// Query runs a query and return the result.
	Query(tx *sql.Tx, query string, parameters map[string]interface{}) ([]map[string]interface{}, error)

	// NewTransaction creates a new database transaction.
	NewTransaction() (*sql.Tx, error)

	// Destruct will close the database adapter releasing all resources.
	Destruct()
}
