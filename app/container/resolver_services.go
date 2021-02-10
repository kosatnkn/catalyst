package container

import (
	"github.com/kosatnkn/catalyst/app/config"
	"github.com/kosatnkn/catalyst/externals/services"
)

var rs Services

// resolveServices resolves all services.
func resolveServices(cfgs []config.ServiceConfig) Services {

	rs.SampleService = services.NewSampleService(getServiceConfigByName(cfgs, "sample"))

	return rs
}
