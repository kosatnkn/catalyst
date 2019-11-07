package adapters

import (
	"context"

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
func (a *PostgresTxAdapter) Wrap(ctx context.Context, fn func(ctx context.Context) (interface{}, error)) (interface{}, error) {

	// TODO: need to create and bind a transaction to context.
	// check to see whether there is a transaction already in the context.
	// this will probably mean that the calling function has been wrapped
	// in a transaction in a previous stage. If so use the same transaction.

	tx, err := a.dba.NewTransaction()
	if err != nil {
		return nil, err
	}

	res, err := fn(ctx)
	if err != nil {

		tx.Rollback()
		return nil, err
	}

	tx.Commit()

	return res, nil
}
