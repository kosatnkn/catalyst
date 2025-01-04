package controllers

import (
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/kosatnkn/catalyst/v3/app/container"
	"github.com/kosatnkn/catalyst/v3/app/transport/http/controllers"
	"github.com/kosatnkn/catalyst/v3/domain/entities"
	"github.com/kosatnkn/catalyst/v3/domain/usecases/sample"
	"github.com/kosatnkn/catalyst/v3/transport/http/request/unpackers"
	"github.com/kosatnkn/catalyst/v3/transport/http/response/transformers"
)

// SampleController contains controller logic for endpoints.
type SampleController struct {
	*controllers.Controller
	sampleUseCase *sample.Sample
}

// NewSampleController creates a new instance of the controller.
func NewSampleController(c *container.Container) *SampleController {
	return &SampleController{
		Controller:    controllers.NewController(c),
		sampleUseCase: sample.NewSample(c),
	}
}

// Get handles retrieving a list of samples.
func (ctl *SampleController) Get(w http.ResponseWriter, r *http.Request) {
	// add a trace string to the request context
	ctx := ctl.WithTrace(r.Context(), "SampleController.Get")

	// get filters from query params
	filters, err := ctl.Filters(r, unpackers.NewSampleFiltersUnpacker())
	if err != nil {
		ctl.SendError(ctx, w, err)
		return
	}

	// get paginator from query parameters
	paginator, err := ctl.Paginator(r)
	if err != nil {
		ctl.SendError(ctx, w, err)
		return
	}

	// get data
	samples, err := ctl.sampleUseCase.Get(ctx, filters, paginator)
	if err != nil {
		ctl.SendError(ctx, w, err)
		return
	}

	// transform school data
	trS, err := ctl.Transform(samples, transformers.NewSampleTransformer(), true)
	if err != nil {
		ctl.SendError(ctx, w, err)
		return
	}

	// transform paginator
	trP, err := ctl.Transform(paginator, transformers.NewPaginatorTransformer(), false)
	if err != nil {
		ctl.SendError(ctx, w, err)
		return
	}

	// send response
	ctl.SendResponse(w, http.StatusOK, trS, trP)
}

// GetByID handles retrieving a single sample.
func (ctl *SampleController) GetByID(w http.ResponseWriter, r *http.Request) {
	// add a trace string to the request context
	ctx := ctl.WithTrace(r.Context(), "SampleController.GetByID")

	// get id from request
	id, _ := strconv.Atoi(ctl.RouteVar(r, "id"))

	// validate
	errs := ctl.Validator.ValidateField("id", id, "required,gt=0")
	if errs != nil {
		ctl.SendError(ctx, w, errs)
		return
	}

	// get data
	sample, err := ctl.sampleUseCase.GetByID(ctx, id)
	if err != nil {
		ctl.SendError(ctx, w, err)
		return
	}

	// transform
	tr, err := ctl.Transform(sample, transformers.NewSampleTransformer(), false)
	if err != nil {
		ctl.SendError(ctx, w, err)
		return
	}

	// send response
	ctl.SendResponse(w, http.StatusOK, tr)
}

// Add adds a new sample entry.
func (ctl *SampleController) Add(w http.ResponseWriter, r *http.Request) {
	// add a trace string to the request context
	ctx := ctl.WithTrace(r.Context(), "SampleController.Add")

	// unpack request
	su := unpackers.NewSampleUnpacker()
	err := ctl.UnpackBody(r, su)
	if err != nil {
		ctl.SendError(ctx, w, err)
		return
	}

	// bind unpacked data to entities
	sample := entities.Sample{
		Name:     su.Name,
		Password: su.Password,
	}

	// add
	err = ctl.sampleUseCase.Add(ctx, sample)
	if err != nil {
		ctl.SendError(ctx, w, err)
		return
	}

	// transform
	// tr := ctl.transform(sample, transformers.NewSampleTransformer(), false)

	// send response
	ctl.SendResponse(w, http.StatusCreated)
}

// Edit updates an existing sample entry.
func (ctl *SampleController) Edit(w http.ResponseWriter, r *http.Request) {
	// add a trace string to the request context
	ctx := ctl.WithTrace(r.Context(), "SampleController.Edit")

	// get id from request
	id, _ := strconv.ParseUint(ctl.RouteVar(r, "id"), 10, 64)

	// validate request parameters
	errs := ctl.Validator.ValidateField("id", id, "required,gt=0")
	if errs != nil {
		ctl.SendError(ctx, w, errs)
		return
	}

	// unpack request body
	su := unpackers.NewSampleUnpacker()
	err := ctl.UnpackBody(r, su)
	if err != nil {
		ctl.SendError(ctx, w, err)
		return
	}

	// bind unpacked data to entities
	sample := entities.Sample{
		ID:       id,
		Name:     su.Name,
		Password: su.Password,
	}

	// edit
	err = ctl.sampleUseCase.Edit(ctx, sample)
	if err != nil {
		ctl.SendError(ctx, w, err)
		return
	}

	// send response
	ctl.SendResponse(w, http.StatusNoContent)
}

// Delete deletes an existing sample entry.
func (ctl *SampleController) Delete(w http.ResponseWriter, r *http.Request) {
	// add a trace string to the request context
	ctx := ctl.WithTrace(r.Context(), "SampleController.Delete")

	// get id from request
	vars := mux.Vars(r)
	id, _ := strconv.Atoi(vars["id"])

	// validate request parameters
	errs := ctl.Validator.ValidateField("id", id, "required,gt=0")
	if errs != nil {
		ctl.SendError(ctx, w, errs)
		return
	}

	// delete
	err := ctl.sampleUseCase.Delete(ctx, id)
	if err != nil {
		ctl.SendError(ctx, w, err)
		return
	}

	// send response
	ctl.SendResponse(w, http.StatusNoContent)
}
