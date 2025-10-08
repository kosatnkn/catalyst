package usecases

import (
	"context"
	"errors"

	"github.com/kosatnkn/catalyst/v3/domain/boundary"
	"github.com/kosatnkn/catalyst/v3/domain/entities"
	"github.com/kosatnkn/catalyst/v3/infra"
)

type AccountUseCases struct {
	retriever boundary.AccountRetriever
	persister boundary.AccountPersister
}

func NewAccountUseCases(c *infra.Container) *AccountUseCases {
	return &AccountUseCases{
		retriever: c.AccountRetriever,
		persister: c.AccountPersister,
	}
}

func (a *AccountUseCases) GetAccounts(ctx context.Context, filters map[string]any, paging map[string]uint32) ([]entities.Account, error) {
	// define allowed filters
	allowedFilterKeys := []string{
		"name",
	}

	return a.retriever.Get(ctx, allowedFiltersOnly(filters, allowedFilterKeys), paging)
}

func (a *AccountUseCases) CreateAccount(ctx context.Context, acc entities.Account) (entities.Account, error) {
	// check whether account with details already exists
	filters := map[string]any{
		"owner": acc.Owner,
	}
	accounts, err := a.retriever.Get(ctx, filters, map[string]uint32{"page": 1, "size": 5})
	if err != nil {
		return entities.Account{}, errors.Join(
			errors.New("usecase-create-account: error retrieving account"),
			err)
	}
	if len(accounts) > 0 {
		return entities.Account{}, errors.New("usecase-create-account: such account already exists")
	}

	// create account
	account, err := a.persister.Create(ctx, acc)
	if err != nil {
		return entities.Account{}, errors.Join(
			errors.New("usecase-create-account: error creating account"),
			err)
	}

	// return details
	return account, nil
}
