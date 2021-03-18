package repositories

import (
	"context"

	"github.com/kosatnkn/catalyst/v2/domain/entities"
	"github.com/kosatnkn/req/filter"
	"github.com/kosatnkn/req/paginator"
)

// SampleRepositoryInterface contract to manipulate `sample` database entity
type SampleRepositoryInterface interface {

	// Get retrieves a collection of samples.
	Get(ctx context.Context, fts []filter.Filter, pgn paginator.Paginator) ([]entities.Sample, error)

	// GetByID retrieves a single sample identified by the id.
	GetByID(ctx context.Context, id int) (entities.Sample, error)

	// Add creates a new sample.
	Add(ctx context.Context, sample entities.Sample) error

	// Edit updates an existing sample identified by the id.
	Edit(ctx context.Context, sample entities.Sample) error

	// Delete deletes an existing sample identified by the id.
	Delete(ctx context.Context, id int) error
}
