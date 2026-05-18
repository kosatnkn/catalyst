package postgres

import (
	"context"

	"github.com/kosatnkn/catalyst/v3/domain/entities"
	"github.com/kosatnkn/catalyst/v3/persistence"
)

type AccountRetrieverPostgres struct {
	db persistence.DatabaseAdapter

	// NOTE: embed the helper to facilitate readiness updates
	*readinessCheckHelper
}

// NewAccountRetrieverPostgres creates a new instance.
func NewAccountRetrieverPostgres(adapter persistence.DatabaseAdapter, ready persistence.DatabaseReadinessAdapter) *AccountRetrieverPostgres {
	return &AccountRetrieverPostgres{
		db:                   adapter,
		readinessCheckHelper: newReadinessCheckHelper(ready),
	}
}

// Get retrieves a slice of accounts that matches the filter.
func (r *AccountRetrieverPostgres) Get(ctx context.Context, filters map[string]any, paging map[string]uint32) ([]entities.Account, error) {
	query := `SELECT * FROM account WHERE name LIKE %?name%`

	params := map[string]any{
		"name": filters["name"],
	}

	// add pagination
	query = withPagination(query, paging)

	result, err := r.db.Query(ctx, query, params)
	if err != nil {
		return nil, r.withReadinessCheck(err) // NOTE: pipe the error returned by the db adapter to update readiness state of the component
	}

	accounts, err := r.mapResult(result)
	if err != nil {
		return nil, err
	}

	return accounts, nil
}

// mapResult to outlets.
//
// NOTE: There are many ways to write mappers. This here is one way of mapping used to demonstrate.
func (r *AccountRetrieverPostgres) mapResult(result []map[string]any) ([]entities.Account, error) {
	accounts := make([]entities.Account, 0, len(result))

	for i, row := range result {
		var account entities.Account
		var err error

		if account.ID, err = to[uint32](row["id"], i); err != nil {
			return accounts, err
		}
		if account.Owner, err = to[string](row["owner_name"], i); err != nil {
			return accounts, err
		}
		if account.Currency, err = to[string](row["bank_name"], i); err != nil {
			return accounts, err
		}
		if account.Balance, err = to[float32](row["balance"], i); err != nil {
			return accounts, err
		}

		accounts = append(accounts, account)
	}

	return accounts, nil
}
