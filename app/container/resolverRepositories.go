package container

import (
	"github.com/kosatnkn/catalyst/externals/repositories"
)

// Resolve all repositories.
func resolveRepositories() Repositories {

	return Repositories{
		SampleRepository: &repositories.SampleRepository{DBAdapter: resolvedAdapters.DB},
	}
}
