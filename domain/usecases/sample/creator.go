package sample

import (
	"context"
	"database/sql"

	"github.com/kosatnkn/catalyst/domain/entities"
)

// Add creates a new sample entry.
func (s *Sample) Add(ctx context.Context, sample entities.Sample) error {

	// business logic here

	_, err := s.transaction.Wrap(func(*sql.Tx) (interface{}, error) {

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
