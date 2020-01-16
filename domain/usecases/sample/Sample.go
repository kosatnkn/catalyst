package sample

import (
	"github.com/kosatnkn/catalyst/app/container"
	"github.com/kosatnkn/catalyst/domain/boundary/adapters"
	"github.com/kosatnkn/catalyst/domain/boundary/repositories"
)

// Sample contains all usecases for samples
type Sample struct {
	transaction      adapters.DBTxAdapterInterface
	sampleRepository repositories.SampleRepositoryInterface
}

// NewSample creates a new instance of sample usecase.
func NewSample(container *container.Container) *Sample {

	return &Sample{
		transaction:      container.Adapters.DBTxAdapter,
		sampleRepository: container.Repositories.SampleRepository,
	}
}
