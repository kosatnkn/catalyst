package container

import (
	"github.com/kosatnkn/catalyst/externals/repositories"
)

var resolvedRepositories Repositories

// resolveRepositories resolve all repositories.
func resolveRepositories() Repositories {

	resolvedRepositories.SampleRepository = repositories.NewSampleRepository(resolvedAdapters.DB)

	return resolvedRepositories
}