package sample

import (
	"github.com/kosatnkn/catalyst/v3/app/container"
	"github.com/kosatnkn/catalyst/v3/domain/boundary/repositories"
	"github.com/kosatnkn/catalyst/v3/externals/adapters"
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
