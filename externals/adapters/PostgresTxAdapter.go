package adapters

import (
	"context"
	"database/sql"

	"github.com/kosatnkn/catalyst/domain/boundary/adapters"
	"github.com/kosatnkn/catalyst/domain/globals"
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

	// attach a transaction to context
	ctx, err := a.attachTx(ctx)
	if err != nil {
		return nil, err
	}

	// get a reference to the attached transaction
	tx := ctx.Value(globals.TxKey).(*sql.Tx)

	res, err := fn(ctx)
	if err != nil {

		tx.Rollback()
		return nil, err
	}

	err = tx.Commit()
	if err != nil {
		return nil, err
	}

	return res, nil
}

// attachTx attaches a database transaction to the context.
//
// This will first check to see whether there is a transaction already in the context.
// Having a transaction already attached to context probably means that the calling function
// has been wrapped in a transaction in a previous stage.
// When this is the case use the existing attached transaction.
// Otherwise create a new transaction and attach.
func (a *PostgresTxAdapter) attachTx(ctx context.Context) (context.Context, error) {

	// check tx altready exists
	tx := ctx.Value(globals.TxKey)
	if tx != nil {
		return ctx, nil
	}

	// attach new tx
	tx, err := a.dba.NewTransaction()
	if err != nil {
		return nil, err
	}

	return context.WithValue(ctx, globals.TxKey, tx), nil
}
