package container

import (
	"github.com/kosatnkn/catalyst/v2/app/config"
)

// Resolve resolves the entire container.
//
// The order of resolution is very important. Low level dependencies need to be resolved before high level dependencies.
// It generally happens in this order.
// 		- Adapters
// 		- Repositories
// 		- Services
func Resolve(cfg *config.Config) *Container {

	ctr := Container{}
	ctr.Adapters = resolveAdapters(cfg)
	ctr.Repositories = resolveRepositories(&ctr.Adapters)
	ctr.Services = resolveServices(cfg.Services)

	return &ctr
}
