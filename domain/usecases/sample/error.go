package sample

import (
	"fmt"

	err "github.com/kosatnkn/catalyst/domain/errors"
)

func (r *Sample) errorNoSample(id int) error {

	return err.NewDomainError(1000, fmt.Sprintf("Sample not found for id %d", id))
}
