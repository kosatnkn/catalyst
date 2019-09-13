package sample

import (
	"context"

	"github.com/kosatnkn/catalyst/domain/entities"
)

// Get returns a list of samples
func (s *Sample) Get(ctx context.Context) ([]entities.Sample, error) {

	// get samples
	samples, err := s.SampleRepository.Get(ctx)
	if err != nil {
		return nil, err
	}

	return samples, nil
}
