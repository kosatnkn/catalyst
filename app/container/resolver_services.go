package container

import (
	"github.com/kosatnkn/catalyst/app/config"
	"github.com/kosatnkn/catalyst/externals/services"
)

// resolveServices resolves all services.
func resolveServices(cfgs []config.ServiceConfig) Services {

	rs := Services{}

	rs.SampleService = services.NewSampleService(getServiceConfigByName(cfgs, "sample"))

	return rs
}
