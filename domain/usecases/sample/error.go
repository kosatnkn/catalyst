package sample

import (
	"fmt"

	domainErr "github.com/kosatnkn/catalyst/domain/error"
)

func (r *Sample) throwNoSampleError(id int) error {

	return (&domainErr.DomainError{}).New("Sample not found",
		1000,
		fmt.Sprintf("No sample found for id %d", id))
}
