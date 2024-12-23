package sample

import (
	"context"

	"github.com/kosatnkn/catalyst/v3/domain/entities"
)

// Add creates a new sample entry.
func (s *Sample) Add(ctx context.Context, sample entities.Sample) error {

	// TODO: your business logic here

	_, err := s.db.WrapInTx(ctx, func(ctx context.Context) (interface{}, error) {
		err := s.sampleRepository.Add(ctx, sample)
		if err != nil {
			return nil, err
		}

		return nil, nil
	})
	if err != nil {
		return err
	}

	return nil
}
