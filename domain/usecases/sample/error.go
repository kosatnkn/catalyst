package sample

import (
	"fmt"

	err "github.com/kosatnkn/catalyst/v3/domain/errors"
)

func (s *Sample) errNoSample(id int) error {
	return err.NewDomainError("1000", fmt.Sprintf("Sample not found for id %d", id), nil)
}

func (s *Sample) errCannotGetData(cause error) error {
	return err.NewDomainError("100", "Cannot get data from repository", cause)
}
