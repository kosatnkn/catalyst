package container

import (
	"fmt"

	"github.com/kosatnkn/catalyst/v2/app/config"
	"github.com/kosatnkn/catalyst/v2/externals/services"
)

// resolveServices resolves all services.
func resolveServices(cfgs []config.ServiceConfig) Services {

	svs := Services{}
	svs.SampleService = services.NewSampleService(getServiceConfig(cfgs, "sample"))

	return svs
}

// getServiceConfig returns the service config by the name of the service.
//
// Will panic if there is no service config found for a given service name.
func getServiceConfig(cfgs []config.ServiceConfig, name string) config.ServiceConfig {

	for i := range cfgs {
		if cfgs[i].Name == name {
			return cfgs[i]
		}
	}

	panic(fmt.Sprintf("Cannot find service configurations for `%s` service", name))
}
