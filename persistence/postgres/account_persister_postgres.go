package postgres

import (
	"context"

	"github.com/kosatnkn/catalyst-pkgs/persistence"
	"github.com/kosatnkn/catalyst/v3/domain/entities"
)

type AccountPersisterPostgres struct {
	db persistence.DatabaseAdapter
}

// NewAccountPersisterPostgres creates a new instance.
func NewAccountPersisterPostgres(adapter persistence.DatabaseAdapter) *AccountPersisterPostgres {
	return &AccountPersisterPostgres{
		db: adapter,
	}
}

// Create creates a new account.
func (p *AccountPersisterPostgres) Create(ctx context.Context, account entities.Account) (entities.Account, error) {
	// TODO: implement
	return entities.Account{}, nil
}

// Update updates an existing account.
func (p *AccountPersisterPostgres) Update(ctx context.Context, account entities.Account) error {
	// TODO: implement
	return nil
}

// Delete deletes an existing account identified by the id.
func (p *AccountPersisterPostgres) Delete(ctx context.Context, id uint32) error {
	// TODO: implement
	return nil
}
