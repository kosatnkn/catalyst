package container

import (
	"github.com/kosatnkn/catalyst/externals/repositories"
)

// Resolve all repositories.
func resolveRepositories() Repositories {

	return Repositories{
		TestRepository: &repositories.TestRepository{DBAdapter: resolvedAdapters.DB},
	}
}
