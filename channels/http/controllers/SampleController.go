package controllers

import (
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/kosatnkn/catalyst/app/container"
	"github.com/kosatnkn/catalyst/channels/http/request"
	"github.com/kosatnkn/catalyst/channels/http/request/unpackers"
	"github.com/kosatnkn/catalyst/channels/http/response"
	"github.com/kosatnkn/catalyst/channels/http/response/transformers"
	"github.com/kosatnkn/catalyst/domain/entities"
	"github.com/kosatnkn/catalyst/domain/usecases/sample"
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

	// get data
	samples, err := ctl.sampleUseCase.Get(ctx)
	if err != nil {
		ctl.sendError(ctx, w, err)
		return
	}

	// transform
	tr, err := response.Transform(samples, transformers.NewSampleTransformer(), true)
	if err != nil {
		ctl.sendError(ctx, w, err)
		return
	}

	// send response
	ctl.sendResponse(ctx, w, http.StatusOK, tr)
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
	tr, err := response.Transform(sample, transformers.NewSampleTransformer(), false)
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
	sampleUnpacker := unpackers.NewSampleUnpacker()
	err := request.Unpack(r, sampleUnpacker)
	if err != nil {
		ctl.sendError(ctx, w, err)
		return
	}

	// validate unpacked data
	errs := ctl.validator.Validate(sampleUnpacker)
	if errs != nil {
		ctl.sendError(ctx, w, errs)
		return
	}

	// bind unpacked data to entities
	sample := entities.Sample{
		Name:     sampleUnpacker.Name,
		Password: sampleUnpacker.Password,
	}

	// add
	err = ctl.sampleUseCase.Add(ctx, sample)
	if err != nil {
		ctl.sendError(ctx, w, err)
		return
	}

	// transform
	// tr := response.Transform(sample, transformers.NewSampleTransformer(), false)

	// send response
	ctl.sendResponse(ctx, w, http.StatusCreated)
}

// Edit updates an existing sample entry.
func (ctl *SampleController) Edit(w http.ResponseWriter, r *http.Request) {

	// get the context
	ctx := r.Context()

	// add a trace string to the context
	ctx = ctl.withTrace(ctx, "SampleController.Edit")

	// unpack request
	sampleUnpacker := unpackers.NewSampleUnpacker()
	err := request.Unpack(r, sampleUnpacker)
	if err != nil {
		ctl.sendError(ctx, w, err)
		return
	}

	// get id from request
	vars := mux.Vars(r)
	id, _ := strconv.Atoi(vars["id"])

	// validate request parameters
	errs := ctl.validator.ValidateField(id, "required,gt=0")
	if errs != nil {
		ctl.sendError(ctx, w, errs)
		return
	}

	// validate unpacked data
	errs = ctl.validator.Validate(sampleUnpacker)
	if errs != nil {
		ctl.sendError(ctx, w, errs)
		return
	}

	// bind unpacked data to entities
	sample := entities.Sample{
		ID:       id,
		Name:     sampleUnpacker.Name,
		Password: sampleUnpacker.Password,
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
