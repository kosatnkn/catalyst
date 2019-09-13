package sample

import "github.com/kosatnkn/catalyst/domain/boundary/repositories"

// Sample contains all usecases for samples
type Sample struct {
	SampleRepository repositories.SampleRepositoryInterface
}
