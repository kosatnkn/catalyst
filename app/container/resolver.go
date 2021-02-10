package container

import (
	"fmt"

	"github.com/kosatnkn/catalyst/app/config"
)

// Resolve resolves the entire container.
//
// The order of resolution is very important. Low level dependencies need to be resolved before high level dependencies.
// It generally happens in this order.
// 		- Adapters
// 		- Repositories
// 		- Services
func Resolve(cfg *config.Config) *Container {

	return &Container{
		Adapters:     resolveAdapters(cfg),
		Repositories: resolveRepositories(),
		Services:     resolveServices(cfg.Services),
	}
}

// getServiceConfigByName returns the service config by name of the service.
func getServiceConfigByName(cfgs []config.ServiceConfig, name string) config.ServiceConfig {

	for i := range cfgs {

		if cfgs[i].Name == name {

			return cfgs[i]
		}
	}

	// must panic if the config is not found
	panic(fmt.Sprintf("Cannot find service configurations for `%s` service", name))
}
