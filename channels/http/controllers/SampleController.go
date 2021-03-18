package controllers

import (
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/kosatnkn/catalyst/v2/app/container"
	"github.com/kosatnkn/catalyst/v2/channels/http/request/unpackers"
	"github.com/kosatnkn/catalyst/v2/channels/http/response/transformers"
	"github.com/kosatnkn/catalyst/v2/domain/entities"
	"github.com/kosatnkn/catalyst/v2/domain/usecases/sample"
)

// SampleController contains controller logic for endpoints.
type SampleController struct {
	*Controller
	sampleUseCase *sample.Sample
}

// NewSampleController creates a new instance of the controller.
func NewSampleController(c *container.Container) *SampleController {

	return &SampleController{
		Controller:    NewController(c),
		sampleUseCase: sample.NewSample(c),
	}
}

// Get handles retreiving a list of samples.
func (ctl *SampleController) Get(w http.ResponseWriter, r *http.Request) {

	// get the context
	ctx := r.Context()

	// add a trace string to the context
	ctx = ctl.withTrace(ctx, "SampleController.Get")

	// get filters from query params
	filters, err := ctl.getFilters(r, unpackers.NewSampleFiltersUnpacker())
	if err != nil {
		ctl.sendError(ctx, w, err)
		return
	}

	// get paginator from query paramaters
	paginator, err := ctl.getPaginator(r)
	if err != nil {
		ctl.sendError(ctx, w, err)
		return
	}

	// get data
	samples, err := ctl.sampleUseCase.Get(ctx, filters, paginator)
	if err != nil {
		ctl.sendError(ctx, w, err)
		return
	}

	// transform school data
	trS, err := ctl.transform(samples, transformers.NewSampleTransformer(), true)
	if err != nil {
		ctl.sendError(ctx, w, err)
		return
	}

	// transform paginator
	trP, err := ctl.transform(paginator, transformers.NewPaginatorTransformer(), false)
	if err != nil {
		ctl.sendError(ctx, w, err)
		return
	}

	// send response
	ctl.sendResponse(ctx, w, http.StatusOK, trS, trP)
}

// GetByID handles retreiving a single sample.
func (ctl *SampleController) GetByID(w http.ResponseWriter, r *http.Request) {

	// get the context
	ctx := r.Context()

	// add a trace string to the context
	ctx = ctl.withTrace(ctx, "SampleController.GetByID")

	// get id from request
	vars := mux.Vars(r)
	id, _ := strconv.Atoi(vars["id"])

	// validate
	errs := ctl.validator.ValidateField(id, "required,gt=0")
	if errs != nil {
		ctl.sendError(ctx, w, errs)
		return
	}

	// get data
	sample, err := ctl.sampleUseCase.GetByID(ctx, id)
	if err != nil {
		ctl.sendError(ctx, w, err)
		return
	}

	// transform
	tr, err := ctl.transform(sample, transformers.NewSampleTransformer(), false)
	if err != nil {
		ctl.sendError(ctx, w, err)
		return
	}

	// send response
	ctl.sendResponse(ctx, w, http.StatusOK, tr)
}

// Add adds a new sample entry.
func (ctl *SampleController) Add(w http.ResponseWriter, r *http.Request) {

	// get the context
	ctx := r.Context()

	// add a trace string to the context
	ctx = ctl.withTrace(ctx, "SampleController.Add")

	// unpack request
	su := unpackers.NewSampleUnpacker()
	err := ctl.unpackBody(r, su)
	if err != nil {
		ctl.sendError(ctx, w, err)
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
		ctl.sendError(ctx, w, err)
		return
	}

	// transform
	// tr := ctl.transform(sample, transformers.NewSampleTransformer(), false)

	// send response
	ctl.sendResponse(ctx, w, http.StatusCreated)
}

// Edit updates an existing sample entry.
func (ctl *SampleController) Edit(w http.ResponseWriter, r *http.Request) {

	// get the context
	ctx := r.Context()

	// add a trace string to the context
	ctx = ctl.withTrace(ctx, "SampleController.Edit")

	// get id from request
	id, _ := strconv.ParseUint(ctl.getRouteVariable(r, "id"), 10, 64)

	// validate request parameters
	errs := ctl.validator.ValidateField(id, "required,gt=0")
	if errs != nil {
		ctl.sendError(ctx, w, errs)
		return
	}

	// unpack request body
	su := unpackers.NewSampleUnpacker()
	err := ctl.unpackBody(r, su)
	if err != nil {
		ctl.sendError(ctx, w, err)
		return
	}

	// bind unpacked data to entities
	sample := entities.Sample{
		ID:       int(id),
		Name:     su.Name,
		Password: su.Password,
	}

	// edit
	err = ctl.sampleUseCase.Edit(ctx, sample)
	if err != nil {
		ctl.sendError(ctx, w, err)
		return
	}

	// send response
	ctl.sendResponse(ctx, w, http.StatusNoContent)
}

// Delete deletes an existing sample entry.
func (ctl *SampleController) Delete(w http.ResponseWriter, r *http.Request) {

	// get the context
	ctx := r.Context()

	// add a trace string to the context
	ctx = ctl.withTrace(ctx, "SampleController.Delete")

	// get id from request
	vars := mux.Vars(r)
	id, _ := strconv.Atoi(vars["id"])

	// validate request parameters
	errs := ctl.validator.ValidateField(id, "required,gt=0")
	if errs != nil {
		ctl.sendError(ctx, w, errs)
		return
	}

	// delete
	err := ctl.sampleUseCase.Delete(ctx, id)
	if err != nil {
		ctl.sendError(ctx, w, err)
		return
	}

	// send response
	ctl.sendResponse(ctx, w, http.StatusNoContent)
}
