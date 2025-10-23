package postgres

import (
	"context"
	"errors"
	"fmt"

	"github.com/kosatnkn/catalyst-pkgs/persistence"
	"github.com/kosatnkn/catalyst/v3/domain/entities"
)

type AccountRetrieverPostgres struct {
	name string
	db   persistence.DatabaseAdapter
}

// NewAccountRetrieverPostgres creates a new instance.
func NewAccountRetrieverPostgres(adapter persistence.DatabaseAdapter) *AccountRetrieverPostgres {
	return &AccountRetrieverPostgres{
		name: "account-retriever-postgres",
		db:   adapter,
	}
}

// Get retrieves a slice of accounts that matches the filter.
func (r *AccountRetrieverPostgres) Get(ctx context.Context, filters map[string]any, paging map[string]uint32) ([]entities.Account, error) {
	// define allowed filters
	allowedFilterKeys := []string{
		"name",
	}
	filters = allowedFiltersOnly(filters, allowedFilterKeys)

	var a []entities.Account

	// TODO: process filters to update both the query and parameters
	query := `SELECT * FROM account WHERE name LIKE ?name`
	params := make(map[string]any)
	params["name"] = filters["name"]

	// add pagination
	query = withPagination(query, paging)

	// DEBUG:
	fmt.Println(query)

	accounts, err := r.db.Query(ctx, query, params)
	if err != nil {
		return a, errors.Join(fmt.Errorf("%s: error retrieving accounts", r.name), err)
	}

	for _, account := range accounts {
		fmt.Println(account)
		a = append(a, entities.Account{
			// TODO: do mapping
		})
	}

	return a, nil
}
