package adapters

import "github.com/kosatnkn/catalyst/app/config"

// DBAdapterInterface is implemeted by all database adapters.
type DBAdapterInterface interface {

	// New creates a new instance of database adapter implementation.
	New(config config.DBConfig) (DBAdapterInterface, error)

	// Query runs a query and return the result.
	Query(query string, parameters map[string]interface{}) ([]map[string]interface{}, error)

	// Destruct will close the database adapter releasing all resources.
	Destruct()
}
