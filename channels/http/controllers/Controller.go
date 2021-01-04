package controllers

import (
	"context"
	"net/http"

	"github.com/kosatnkn/catalyst/app/container"
	"github.com/kosatnkn/catalyst/channels/http/response"
	"github.com/kosatnkn/catalyst/domain/boundary/adapters"
	"github.com/kosatnkn/catalyst/domain/globals"
)

// Controller is the base struct that holds fields and functionality common to all controllers.
type Controller struct {
	logger    adapters.LogAdapterInterface
	validator adapters.ValidatorAdapterInterface
}

// NewController creates a new instance of the controller.
func NewController(container *container.Container) *Controller {

	return &Controller{
		logger:    container.Adapters.LogAdapter,
		validator: container.Adapters.ValidatorAdapter,
	}
}

// withTrace adds an optional tracing string that will be displayed in error messages.
func (ctl *Controller) withTrace(ctx context.Context, prefix string) context.Context {
	return globals.AddTrace(ctx, prefix)
}

// sendResponse is a convenience function wrapping the actual `response.Send` function
// to provide a cleaner usage interface.
func (ctl *Controller) sendResponse(ctx context.Context, w http.ResponseWriter, code int, payload ...interface{}) {

	if len(payload) == 0 {
		response.Send(w, nil, code)
		return
	}

	response.Send(w, response.Map(payload), code)
}

// sendError is a convenience function wrapping the actual `response.Error` function
// to provide a cleaner usage interface.
func (ctl *Controller) sendError(ctx context.Context, w http.ResponseWriter, err interface{}) {
	response.Error(ctx, w, err, ctl.logger)
}
