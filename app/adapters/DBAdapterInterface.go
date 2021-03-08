package adapters

import "github.com/kosatnkn/db"

// DBAdapterInterface is implemented by all database adapters.
type DBAdapterInterface interface {
	db.AdapterInterface
}
