package usecases

import (
	"context"
	"errors"
	"fmt"

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
	return a.retriever.Get(ctx, filters, paging)
}

func (a *AccountUseCases) CreateAccount(ctx context.Context, acc entities.Account) (entities.Account, error) {
	ident := "usecase-create-account"

	// check whether account with details already exists
	filters := map[string]any{
		"owner": acc.Owner,
	}
	accounts, err := a.retriever.Get(ctx, filters, map[string]uint32{"page": 1, "size": 5})
	if err != nil {
		return entities.Account{}, errors.Join(fmt.Errorf("%s: error retrieving account", ident), err)
	}
	if len(accounts) > 0 {
		return entities.Account{}, fmt.Errorf("%s: such account already exists", ident)
	}

	// create account
	account, err := a.persister.Create(ctx, acc)
	if err != nil {
		return entities.Account{}, errors.Join(fmt.Errorf("%s: error creating account", ident), err)
	}

	// return details
	return account, nil
}
