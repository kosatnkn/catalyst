package container

import (
	"github.com/kosatnkn/catalyst/v2/app/config"
	"github.com/kosatnkn/catalyst/v2/externals/services"
)

// resolveServices resolves all services.
func resolveServices(cfgs []config.ServiceConfig) Services {

	svs := Services{}

	svs.SampleService = services.NewSampleService(getServiceConfigByName(cfgs, "sample"))

	return svs
}
