package container

import (
	"fmt"

	"github.com/kosatnkn/catalyst/app/config"
	"github.com/kosatnkn/catalyst/externals/services"
)

var resolvedServices Services
var serviceConfigs []config.ServiceConfig

// resolveServices resolves all services.
func resolveServices(cfgs []config.ServiceConfig) Services {

	serviceConfigs = cfgs

	resolvedServices.SampleService = services.NewSampleService(getConfigByName("sample"))

	return resolvedServices
}

// getConfigByName returns the service config by name of the service
func getConfigByName(name string) config.ServiceConfig {

	for i := range serviceConfigs {

		if serviceConfigs[i].Name == name {

			return serviceConfigs[i]
		}
	}

	// must panic if the config is not found
	panic(fmt.Sprintf("Cannot find service configurations for `%s` service", name))
}
