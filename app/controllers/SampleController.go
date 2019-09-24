package controllers

import (
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/kosatnkn/catalyst/app/container"
	"github.com/kosatnkn/catalyst/app/error"
	"github.com/kosatnkn/catalyst/app/transport/request"
	"github.com/kosatnkn/catalyst/app/transport/request/unpackers"
	"github.com/kosatnkn/catalyst/app/transport/response"
	"github.com/kosatnkn/catalyst/app/transport/response/transformers"
	"github.com/kosatnkn/catalyst/domain/entities"
	"github.com/kosatnkn/catalyst/domain/globals"
	"github.com/kosatnkn/catalyst/domain/usecases/sample"
)

// SampleController contains controller logic for endpoints.
type SampleController struct {
	container     *container.Container
	sampleUseCase *sample.Sample
}

// NewSampleController returns a base type for this controller
func NewSampleController(container *container.Container) *SampleController {

	// create instance of controller
	return &SampleController{
		container:     container,
		sampleUseCase: sample.NewSample(container),
	}
}

// Get handles retreiving a list of samples.
func (ctl *SampleController) Get(w http.ResponseWriter, r *http.Request) {

	// get the context
	ctx := r.Context()

	// append a prefix value to the context passed within the request
	ctx = globals.AppendToContextPrefix(ctx, "SampleController.Get")

	// get data
	samples, err := ctl.sampleUseCase.Get(ctx)
	if err != nil {
		error.Handle(ctx, err, w, ctl.container.Adapters.Log)
		return
	}

	// transform
	tr := response.Transform(samples, transformers.NewSampleTransformer(), true)

	// send response
	response.Send(w, tr, http.StatusOK)
}

// GetByID handles retreiving a single sample.
func (ctl *SampleController) GetByID(w http.ResponseWriter, r *http.Request) {

	// get the context
	ctx := r.Context()

	// append a prefix value to the context passed within the request
	ctx = globals.AppendToContextPrefix(ctx, "SampleController.GetByID")

	// get id from request
	vars := mux.Vars(r)
	id, _ := strconv.Atoi(vars["id"])

	// validate
	// NOTE: here a validation is not actually needed since query parameters
	// are validated to a certain extent by putting parameter validations in
	// routes and by data type conversions done in the controller
	errs := ctl.container.Adapters.Validator.ValidateField(id, "required,gt=0")
	if errs != nil {
		error.HandleValidationErrors(ctx, errs, w, ctl.container.Adapters.Log)
		return
	}

	// get data
	sample, err := ctl.sampleUseCase.GetByID(ctx, id)
	if err != nil {
		error.Handle(ctx, err, w, ctl.container.Adapters.Log)
		return
	}

	// transform
	tr := response.Transform(sample, transformers.NewSampleTransformer(), false)

	// send response
	response.Send(w, tr, http.StatusOK)
}

// Add adds a new sample entry.
func (ctl *SampleController) Add(w http.ResponseWriter, r *http.Request) {

	// get the context
	ctx := r.Context()

	// append a prefix value to the context passed within the request
	ctx = globals.AppendToContextPrefix(ctx, "SampleController.Add")

	// unpack request
	sampleUnpacker := unpackers.NewSampleUnpacker()
	err := request.Unpack(r, sampleUnpacker)
	if err != nil {
		error.Handle(ctx, err, w, ctl.container.Adapters.Log)
		return
	}

	// validate unpacked data
	errs := ctl.container.Adapters.Validator.Validate(sampleUnpacker)
	if errs != nil {
		error.HandleValidationErrors(ctx, errs, w, ctl.container.Adapters.Log)
		return
	}

	// bind unpacked data to entities
	sample := entities.Sample{
		Name:     sampleUnpacker.Street,
		Password: sampleUnpacker.Planet,
	}

	// add
	err = ctl.sampleUseCase.Add(ctx, sample)
	if err != nil {
		error.Handle(ctx, err, w, ctl.container.Adapters.Log)
		return
	}

	// transform
	// tr := response.Transform(sample, transformers.NewSampleTransformer(), false)

	// send response
	response.Send(w, nil, http.StatusCreated)
}

// Edit updates an existing sample entry.
func (ctl *SampleController) Edit(w http.ResponseWriter, r *http.Request) {

	// get the context
	ctx := r.Context()

	// append a prefix value to the context passed within the request
	ctx = globals.AppendToContextPrefix(ctx, "SampleController.Edit")

	// unpack request
	sampleUnpacker := unpackers.NewSampleUnpacker()
	err := request.Unpack(r, sampleUnpacker)
	if err != nil {
		error.Handle(ctx, err, w, ctl.container.Adapters.Log)
		return
	}

	// get id from request
	vars := mux.Vars(r)
	id, _ := strconv.Atoi(vars["id"])

	// validate request parameters
	errs := ctl.container.Adapters.Validator.ValidateField(id, "required,gt=0")
	if errs != nil {
		error.HandleValidationErrors(ctx, errs, w, ctl.container.Adapters.Log)
		return
	}

	// validate unpacked data
	errs = ctl.container.Adapters.Validator.Validate(sampleUnpacker)
	if errs != nil {
		error.HandleValidationErrors(ctx, errs, w, ctl.container.Adapters.Log)
		return
	}

	// bind unpacked data to entities
	sample := entities.Sample{
		Name:     sampleUnpacker.Street,
		Password: sampleUnpacker.Planet,
	}

	// edit
	err = ctl.sampleUseCase.Edit(ctx, id, sample)
	if err != nil {
		error.Handle(ctx, err, w, ctl.container.Adapters.Log)
		return
	}

	// send response
	response.Send(w, nil, http.StatusNoContent)
}

// Delete deletes an existing sample entry.
func (ctl *SampleController) Delete(w http.ResponseWriter, r *http.Request) {

	// get the context
	ctx := r.Context()

	// append a prefix value to the context passed within the request
	ctx = globals.AppendToContextPrefix(ctx, "SampleController.Delete")

	// get id from request
	vars := mux.Vars(r)
	id, _ := strconv.Atoi(vars["id"])

	// validate request parameters
	errs := ctl.container.Adapters.Validator.ValidateField(id, "required,gt=0")
	if errs != nil {
		error.HandleValidationErrors(ctx, errs, w, ctl.container.Adapters.Log)
		return
	}

	// delete
	err := ctl.sampleUseCase.Delete(ctx, id)
	if err != nil {
		error.Handle(ctx, err, w, ctl.container.Adapters.Log)
		return
	}

	// send response
	response.Send(w, nil, http.StatusNoContent)
}
