package controllers

import (
	"context"
	"io"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/kosatnkn/catalyst/v3/app/adapters"
	"github.com/kosatnkn/catalyst/v3/app/container"
	"github.com/kosatnkn/catalyst/v3/app/transport/http/request"
	"github.com/kosatnkn/catalyst/v3/app/transport/http/request/unpacker"
	"github.com/kosatnkn/catalyst/v3/internal/req/filter"
	"github.com/kosatnkn/catalyst/v3/internal/req/paginator"
	"github.com/kosatnkn/catalyst/v3/transport/http/response"
	"github.com/kosatnkn/catalyst/v3/transport/http/response/transformers"
)

// Controller is the base struct that holds fields and functionality common to all controllers.
type Controller struct {
	Logger    adapters.LogAdapterInterface
	Validator adapters.ValidatorAdapterInterface
	*filter.FilterControllerFacilitator
	*paginator.PaginatorControllerFacilitator
}

// NewController creates a new instance of the controller.
func NewController(c *container.Container) *Controller {
	return &Controller{
		Logger:                         c.Adapters.Log,
		Validator:                      c.Adapters.Validator,
		FilterControllerFacilitator:    filter.NewFilterControllerFacilitator(),
		PaginatorControllerFacilitator: paginator.NewPaginatorControllerFacilitator(),
	}
}

// WithTrace adds an optional tracing string that will be displayed in error messages.
func (ctl *Controller) WithTrace(ctx context.Context, point string) context.Context {
	return ctl.Logger.AppendTracePoint(ctx, point)
}

// RouteVar returns the value of the route variable denoted by the name.
func (ctl *Controller) RouteVar(r *http.Request, name string) string {
	return mux.Vars(r)[name]
}

// filters extracts filters from query parameters.
func (ctl *Controller) Filters(r *http.Request, fu unpacker.UnpackerInterface) ([]filter.Filter, interface{}) {
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
	errs := ctl.Validator.Validate(fu)
	if errs != nil {
		return filters, errs
	}

	return ctl.GetFilters(fu)
}

// Paginator extracts pagination data from query parameters
func (ctl *Controller) Paginator(r *http.Request) (paginator.Paginator, interface{}) {
	// create default paginator
	p := ctl.GetPaginator(1, 10)

	// get paginator from query params
	pgn := r.URL.Query()["paginator"]
	if len(pgn) == 0 {
		return p, nil
	}

	// unpack pagination data sent in query
	pu := paginator.NewPaginatorUnpacker()

	err := request.Unpack([]byte(pgn[0]), pu)
	if err != nil {
		return p, err
	}

	// validate unpacked data
	errs := ctl.Validator.Validate(pu)
	if errs != nil {
		return p, errs
	}

	// bind unpacked data to paginator
	p.Page = pu.Page
	p.Size = uint32(pu.Size)

	return p, nil
}

// UnpackBody unpacks and validates the request body.
func (ctl *Controller) UnpackBody(r *http.Request, u unpacker.UnpackerInterface) interface{} {
	body, err := io.ReadAll(r.Body)
	if err != nil {
		return err
	}

	err = request.Unpack(body, u)
	if err != nil {
		return err
	}

	// validate unpacked data
	errs := ctl.Validator.Validate(u)
	if errs != nil {
		return errs
	}

	return nil
}

// Transform is a convenience function wrapping the actual `response.Transform` function
// to provide a cleaner usage interface.
func (ctl *Controller) Transform(data interface{}, t transformers.TransformerInterface, isCollection bool) (interface{}, error) {
	return response.Transform(data, t, isCollection)
}

// SendResponse is a convenience function wrapping the actual `response.Send` function
// to provide a cleaner usage interface.
func (ctl *Controller) SendResponse(w http.ResponseWriter, code int, payload ...interface{}) {
	if len(payload) == 0 {
		response.Send(w, code, nil)
		return
	}

	response.Send(w, code, payload)
}

// SendError is a convenience function wrapping the actual `response.Error` function
// to provide a cleaner usage interface.
func (ctl *Controller) SendError(ctx context.Context, w http.ResponseWriter, err interface{}) {
	response.Error(ctx, w, ctl.Logger, err)
}
