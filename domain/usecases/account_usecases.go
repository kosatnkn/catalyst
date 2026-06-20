package usecases

import (
	"context"
	"fmt"

	"github.com/kosatnkn/catalyst/v3/domain/boundary"
	"github.com/kosatnkn/catalyst/v3/domain/entities"
	"github.com/kosatnkn/catalyst/v3/infra"
)

// AccountUseCases
type AccountUseCases struct {
	// transaction
	// NOTE: if you are attaching the 'Readiness' probe and uses the database transaction wrapper (tx.WrapInTx)
	// in the domain layer, you must also pass in the 'Readiness' in to the usecase as well.
	// This will catch and evaluate errors thrown from the transaction wrapper in order to set the db state
	// in 'Readiness'. Use the provided 'withDBReadinessCheck' helper function for that.
	tx    boundary.DatabaseTx
	ready boundary.Readiness
	// account
	retriever boundary.AccountRetriever
	persister boundary.AccountPersister
}

// NewAccountUseCases creates a new instance.
func NewAccountUseCases(c *infra.Container) *AccountUseCases {
	return &AccountUseCases{
		tx:        c.DBAdapter,
		ready:     c.Readiness,
		retriever: c.AccountRetriever,
		persister: c.AccountPersister,
	}
}

// GetAccounts filtered and paged.
func (u *AccountUseCases) GetAccounts(ctx context.Context, filters map[string]any, paging map[string]uint32) ([]entities.Account, error) {
	return u.retriever.Get(ctx, filters, paging)
}

// CreateAccount from given data.
func (u *AccountUseCases) CreateAccount(ctx context.Context, acc entities.Account) (entities.Account, error) {
	// execute in a single transaction
	a, txErr := u.tx.WrapInTx(ctx, func(ctx context.Context) (any, error) {
		// check whether account with details already exists
		filters := map[string]any{
			"owner": acc.Owner,
		}
		accounts, err := u.retriever.Get(ctx, filters, map[string]uint32{"page": 1, "size": 1})
		if err != nil {
			return entities.Account{}, err
		}
		if len(accounts) > 0 {
			return entities.Account{}, fmt.Errorf("account already exists")
		}

		// create account
		return u.persister.Create(ctx, acc)
	})

	return a.(entities.Account), withDBReadinessCheck(u.tx, u.ready, txErr)
}
