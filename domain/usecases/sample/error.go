package sample

import (
	"fmt"

	err "github.com/kosatnkn/catalyst/domain/errors"
)

func (r *Sample) errorNoSample(id int) error {

	return err.NewDomainError("Sample not found",
		1000,
		fmt.Sprintf("No sample found for id %d", id))
}
