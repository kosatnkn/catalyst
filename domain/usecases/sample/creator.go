package sample

import (
	"context"

	"github.com/kosatnkn/catalyst/domain/entities"
)

// Add creates a new sample entry.
func (s *Sample) Add(ctx context.Context, sample entities.Sample) error {

	// business logic here

	err := s.sampleRepository.Add(ctx, sample)
	if err != nil {
		return err
	}

	return nil
}
