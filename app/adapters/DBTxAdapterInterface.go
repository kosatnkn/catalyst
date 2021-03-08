package adapters

import "github.com/kosatnkn/db"

// DBTxAdapterInterface is implemented by all database transaction adapters.
type DBTxAdapterInterface interface {
	db.TxAdapterInterface
}
