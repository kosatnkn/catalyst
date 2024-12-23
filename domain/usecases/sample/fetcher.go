package sample

import (
	"context"

	"github.com/kosatnkn/catalyst/v3/domain/entities"
	"github.com/kosatnkn/req/filter"
	"github.com/kosatnkn/req/paginator"
)

// Get returns a list of samples.
func (s *Sample) Get(ctx context.Context, fts []filter.Filter, pgn paginator.Paginator) ([]entities.Sample, error) {
	// get samples
	samples, err := s.sampleRepository.Get(ctx, fts, pgn)
	if err != nil {
		return nil, s.errCannotGetData(err)
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
		return entities.Sample{}, s.errNoSample(id)
	}

	return sample, nil
}
