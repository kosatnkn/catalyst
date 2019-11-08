package repositories

import (
	"context"

	"github.com/kosatnkn/catalyst/domain/entities"
)

// SampleRepositoryInterface contract to manipulate `sample` database entity
type SampleRepositoryInterface interface {

	// Get retrieves a collection of Samples.
	Get(ctx context.Context) ([]entities.Sample, error)

	// GetByID retrieves a single Sample.
	GetByID(ctx context.Context, id int) (entities.Sample, error)

	// Add adds a new sample.
	Add(ctx context.Context, sample entities.Sample) error

	// Edit updates an existing sample identified by the id.
	Edit(ctx context.Context, sample entities.Sample) error

	// Delete deletes an existing sample identified by id.
	Delete(ctx context.Context, id int) error
}
