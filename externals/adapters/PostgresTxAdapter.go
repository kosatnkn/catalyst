package adapters

import (
	"database/sql"

	"github.com/kosatnkn/catalyst/domain/boundary/adapters"
)

// PostgresTxAdapter is used to handle postgres db transactions.
type PostgresTxAdapter struct {
	dba adapters.DBAdapterInterface
}

// NewPostgresTxAdapter creates a new Postgres transaction adapter instance.
func NewPostgresTxAdapter(dba adapters.DBAdapterInterface) adapters.DBTxAdapterInterface {

	return &PostgresTxAdapter{
		dba: dba,
	}
}

// Wrap runs the content of the function in a single transaction.
func (a *PostgresTxAdapter) Wrap(fn func(*sql.Tx) (interface{}, error)) (interface{}, error) {

	tx, err := a.dba.NewTransaction()
	if err != nil {
		return nil, err
	}

	res, err := fn(tx)
	if err != nil {

		tx.Rollback()
		return nil, err
	}

	tx.Commit()

	return res, nil
}
