package container

import (
	"github.com/kosatnkn/catalyst/externals/repositories"
)

var rr Repositories

// resolveRepositories resolve all repositories.
func resolveRepositories() Repositories {

	rr.SampleRepository = repositories.NewSampleMySQLRepository(ra.DBAdapter)

	return rr
}
