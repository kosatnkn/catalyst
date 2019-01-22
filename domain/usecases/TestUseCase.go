package usecases

import (
	"github.com/kosatnkn/catalyst/domain/boundary/repositories"
	"github.com/kosatnkn/catalyst/domain/entities"
)

type TestUseCase struct {
	TestRepository repositories.TestRepositoryInterface
}

// Test databse functionality.
func (t *TestUseCase) TestDB() ([]entities.Test, error) {
	return t.TestRepository.TestRepositoryMethod()
}
