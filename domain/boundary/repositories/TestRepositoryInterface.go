package repositories

import "github.com/kosatnkn/catalyst/domain/entities"

type TestRepositoryInterface interface {
	TestRepositoryMethod() ([]entities.Test, error)
}
