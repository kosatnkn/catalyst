package container

import (
	"github.com/kosatnkn/catalyst/v3/externals/repositories"
)

// resolveRepositories resolve all repositories.
func resolveRepositories(ats *Adapters) Repositories {
	rts := Repositories{}
	rts.SampleRepository = repositories.NewSamplePostgresRepository(ats.DB)

	return rts
}
