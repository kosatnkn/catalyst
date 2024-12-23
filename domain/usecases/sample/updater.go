package sample

import (
	"context"

	"github.com/kosatnkn/catalyst/v3/domain/entities"
)

// Edit updates an existing sample entry.
func (s *Sample) Edit(ctx context.Context, sample entities.Sample) error {

	// TODO: your business logic here

	err := s.sampleRepository.Edit(ctx, sample)
	if err != nil {
		return err
	}

	return nil
}
