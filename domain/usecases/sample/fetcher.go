package sample

import (
	"context"

	"github.com/kosatnkn/catalyst/domain/entities"
)

// Get returns a list of samples.
func (s *Sample) Get(ctx context.Context) ([]entities.Sample, error) {

	// get samples
	samples, err := s.sampleRepository.Get(ctx)
	if err != nil {
		return nil, err
	}

	return samples, nil
}

// GetByID returns a single sample object by id.
func (s *Sample) GetByID(ctx context.Context, id int) (entities.Sample, error) {

	// get sample
	sample, err := s.sampleRepository.GetByID(ctx, id)
	if err != nil {
		return entities.Sample{}, err
	}

	if sample.ID == 0 {
		return entities.Sample{}, s.errorNoSample(id)
	}

	return sample, nil
}
