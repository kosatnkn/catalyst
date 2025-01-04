package container

import (
	"github.com/kosatnkn/catalyst/v3/externals/repositories/postgres"
)

// resolveRepositories resolve all repositories.
func resolveRepositories(ats *Adapters) Repositories {
	rts := Repositories{}
	rts.SampleRepository = postgres.NewSamplePostgresRepository(ats.DB)

	return rts
}
