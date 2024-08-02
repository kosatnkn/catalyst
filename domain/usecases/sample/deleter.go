package sample

import (
	"context"
)

// Delete removes a sample entry.
func (s *Sample) Delete(ctx context.Context, id int) error {

	// TODO: your business logic here

	err := s.sampleRepository.Delete(ctx, id)
	if err != nil {
		return err
	}

	return nil
}
