package repositories

import (
	"context"

	"github.com/kosatnkn/catalyst/domain/entities"
)

// SampleRepositoryInterface contract to manipulate `sample` database entity
type SampleRepositoryInterface interface {

	// Get returns a slice of Samples
	Get(ctx context.Context) ([]entities.Sample, error)
}
