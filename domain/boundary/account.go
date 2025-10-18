package boundary

import (
	"context"

	"github.com/kosatnkn/catalyst/v3/domain/entities"
)

type AccountPersister interface {
	// Create creates a new account.
	Create(ctx context.Context, account entities.Account) (entities.Account, error)

	// Update updates an existing account.
	Update(ctx context.Context, account entities.Account) error

	// Delete deletes an existing account identified by the id.
	Delete(ctx context.Context, id uint32) error
}

type AccountRetriever interface {
	// Get retrieves a slice of accounts that matches the filter.
	Get(ctx context.Context, filter map[string]any, paging map[string]uint32) ([]entities.Account, error)
}
