package sample

import (
	"context"
)

// Delete removes a sample entry.
func (s *Sample) Delete(ctx context.Context, id int) error {

	// business logic here

	err := s.SampleRepository.Delete(ctx, id)
	if err != nil {
		return err
	}

	return nil
}
