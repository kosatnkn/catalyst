package sample

import (
	"github.com/kosatnkn/catalyst/v2/app/adapters"
	"github.com/kosatnkn/catalyst/v2/app/container"
	"github.com/kosatnkn/catalyst/v2/domain/boundary/repositories"
)

// Sample contains all usecases for samples
type Sample struct {
	db               adapters.DBAdapterInterface
	sampleRepository repositories.SampleRepositoryInterface
}

// NewSample creates a new instance of sample usecase.
func NewSample(ctr *container.Container) *Sample {
	return &Sample{
		db:               ctr.Adapters.DB,
		sampleRepository: ctr.Repositories.SampleRepository,
	}
}
