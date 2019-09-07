package adapters

// DBAdapterInterface is implemeted by all database adapters.
type DBAdapterInterface interface {

	// Query runs a query and return the result.
	Query(query string, parameters map[string]interface{}) ([]map[string]interface{}, error)

	// Destruct will close the database adapter releasing all resources.
	Destruct()
}
