package controllers

import (
	"context"
	"io"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/kosatnkn/catalyst/v2/app/adapters"
	"github.com/kosatnkn/catalyst/v2/app/container"
	"github.com/kosatnkn/catalyst/v2/transport/http/request"
	"github.com/kosatnkn/catalyst/v2/transport/http/request/unpackers"
	"github.com/kosatnkn/catalyst/v2/transport/http/response"
	"github.com/kosatnkn/catalyst/v2/transport/http/response/transformers"
	"github.com/kosatnkn/req/filter"
	"github.com/kosatnkn/req/paginator"
)

// Controller is the base struct that holds fields and functionality common to all controllers.
type Controller struct {
	logger    adapters.LogAdapterInterface
	validator adapters.ValidatorAdapterInterface
	*filter.FilterControllerFacilitator
	*paginator.PaginatorControllerFacilitator
}

// NewController creates a new instance of the controller.
func NewController(c *container.Container) *Controller {
	return &Controller{
		logger:                         c.Adapters.Log,
		validator:                      c.Adapters.Validator,
		FilterControllerFacilitator:    filter.NewFilterControllerFacilitator(),
		PaginatorControllerFacilitator: paginator.NewPaginatorControllerFacilitator(),
	}
}

// withTrace adds an optional tracing string that will be displayed in error messages.
func (ctl *Controller) withTrace(ctx context.Context, point string) context.Context {
	return ctl.logger.AppendTracePoint(ctx, point)
}

// routeVar returns the value of the route variable denoted by the name.
func (ctl *Controller) routeVar(r *http.Request, name string) string {
	return mux.Vars(r)[name]
}

// filters extracts filters from query parameters.
func (ctl *Controller) filters(r *http.Request, fu unpackers.UnpackerInterface) ([]filter.Filter, interface{}) {
	// create empty filters slice
	filters := make([]filter.Filter, 0)

	// get paginator from query params
	fts := r.URL.Query()["filters"]
	if len(fts) == 0 {
		return filters, nil
	}

	// unpack filter data sent in query
	err := request.Unpack([]byte(fts[0]), fu)
	if err != nil {
		return filters, err
	}

	// validate unpacked data
	errs := ctl.validator.Validate(fu)
	if errs != nil {
		return filters, errs
	}

	return ctl.GetFilters(fu)
}

// paginator extracts pagination data from query parameters
func (ctl *Controller) paginator(r *http.Request) (paginator.Paginator, interface{}) {
	// create default paginator
	paginator := ctl.GetPaginator(1, 10)

	// get paginator from query params
	pgn := r.URL.Query()["paginator"]
	if len(pgn) == 0 {
		return paginator, nil
	}

	// unpack pagination data sent in query
	pu := unpackers.NewPaginatorUnpacker()

	err := request.Unpack([]byte(pgn[0]), pu)
	if err != nil {
		return paginator, err
	}

	// validate unpacked data
	errs := ctl.validator.Validate(pu)
	if errs != nil {
		return paginator, errs
	}

	// bind unpacked data to paginator
	paginator.Page = pu.Page
	paginator.Size = uint32(pu.Size)

	return paginator, nil
}

// unpackBody unpacks and validates the request body.
func (ctl *Controller) unpackBody(r *http.Request, u unpackers.UnpackerInterface) interface{} {
	body, err := io.ReadAll(r.Body)
	if err != nil {
		return err
	}

	err = request.Unpack(body, u)
	if err != nil {
		return err
	}

	// validate unpacked data
	errs := ctl.validator.Validate(u)
	if errs != nil {
		return errs
	}

	return nil
}

// transform is a convenience function wrapping the actual `response.Transform` function
// to provide a cleaner usage interface.
func (ctl *Controller) transform(data interface{}, t transformers.TransformerInterface, isCollection bool) (interface{}, error) {
	return response.Transform(data, t, isCollection)
}

// sendResponse is a convenience function wrapping the actual `response.Send` function
// to provide a cleaner usage interface.
func (ctl *Controller) sendResponse(ctx context.Context, w http.ResponseWriter, code int, payload ...interface{}) {
	if len(payload) == 0 {
		response.Send(w, code, nil)
		return
	}

	response.Send(w, code, payload)
}

// sendError is a convenience function wrapping the actual `response.Error` function
// to provide a cleaner usage interface.
func (ctl *Controller) sendError(ctx context.Context, w http.ResponseWriter, err interface{}) {
	response.Error(ctx, w, ctl.logger, err)
}
