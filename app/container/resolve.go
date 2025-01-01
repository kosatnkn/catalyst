package container

import (
	"github.com/kosatnkn/catalyst/v3/app/config"
)

// Resolve resolves the entire container.
//
// The order of resolution is very important. Low level dependencies need to be resolved before high level dependencies.
// It generally happens in this order.
//   - Adapters
//   - Repositories
func Resolve(cfg *config.Config) *Container {
	ctr := Container{}
	ctr.Adapters = resolveAdapters(cfg)
	ctr.Repositories = resolveRepositories(&ctr.Adapters)

	return &ctr
}
