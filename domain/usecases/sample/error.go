package sample

import (
	"fmt"

	err "github.com/kosatnkn/catalyst/v3/domain/errors"
)

func (s *Sample) errNoSample(id int) error {
	return err.NewDomainError("sample-usecase-noid", fmt.Sprintf("sample-usecase: sample not found for id %d", id), nil)
}

func (s *Sample) errCannotGetData(cause error) error {
	return err.NewDomainError("sample-usecase-nodata", "sample-usecase: cannot get data from repository", cause)
}
