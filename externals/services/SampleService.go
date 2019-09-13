package services

import (
	"github.com/kosatnkn/catalyst/app/config"
	"github.com/kosatnkn/catalyst/domain/boundary/services"
)

// SampleService is an example service to call a third party web API
type SampleService struct {
	cfg config.ServiceConfig
}

// NewSampleService creates a new instance of the service
func NewSampleService(cfg config.ServiceConfig) services.SampleServiceInterface {

	return &SampleService{cfg: cfg}
}

// SampleServiceMethod sample method
func (svc *SampleService) SampleServiceMethod() {

}
