package container

import (
	"github.com/kosatnkn/catalyst/externals/repositories"
)

// resolveRepositories resolve all repositories.
func resolveRepositories(ats *Adapters) Repositories {

	rts := Repositories{}

	rts.SampleRepository = repositories.NewSampleMySQLRepository(ats.DB)

	return rts
}
