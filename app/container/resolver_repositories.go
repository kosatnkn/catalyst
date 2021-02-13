package container

import (
	"github.com/kosatnkn/catalyst/externals/repositories"
)

// resolveRepositories resolve all repositories.
func resolveRepositories(ats *Adapters) Repositories {

	rr := Repositories{}

	rr.SampleRepository = repositories.NewSampleMySQLRepository(ats.DB)

	return rr
}
