package controllers

import (
	"net/http"

	"github.com/kosatnkn/catalyst/app/container"
	"github.com/kosatnkn/catalyst/app/error"
	"github.com/kosatnkn/catalyst/app/transport/response"
	"github.com/kosatnkn/catalyst/app/transport/response/transformers"
	"github.com/kosatnkn/catalyst/domain/globals"
	"github.com/kosatnkn/catalyst/domain/usecases/sample"
)

// SampleController contains controller logic for endpoints.
type SampleController struct {
	Container     *container.Container
	SampleUseCase sample.Sample
}

// NewSampleController returns a base type for this controller
func NewSampleController(container *container.Container) *SampleController {

	// init use cases
	sampleUseCase := sample.Sample{
		SampleRepository: container.Repositories.SampleRepository,
	}

	// create instance of controller
	return &SampleController{
		Container:     container,
		SampleUseCase: sampleUseCase,
	}
}

// Get handles retreiving a list of samples
func (ctl *SampleController) Get(w http.ResponseWriter, r *http.Request) {

	// get the context
	ctx := r.Context()

	// append a prefix value to the context passed within the request
	ctx = globals.AppendToContextPrefix(ctx, "SampleController.Get")

	// get data
	samples, err := ctl.SampleUseCase.Get(ctx)
	if err != nil {
		error.Handle(ctx, err, w, ctl.Container.Adapters.Log)
		return
	}

	// TODO: add this method to transformer package
	// map result to transformer
	var tr []transformers.SampleTransformer
	for _, sample := range samples {
		tr = append(tr, transformers.SampleTransformer{
			ID:   sample.ID,
			Name: sample.Name,
		})
	}

	// send response
	response.Send(w, response.Transform(tr), http.StatusOK)
}

// // GetByID handles retreiving a single sample
// func (ctl *SampleController) GetByID(w http.ResponseWriter, r *http.Request) {

// 	// append a prefix value to the context passed within the request
// 	r = r.WithContext(globals.AppendToContextPrefix(r.Context(), "SampleController.GetAll"))

// 	// unpack request
// 	sampleUnpacker := unpackers.SampleUnpacker{}
// 	err := request.Unpack(r, &sampleUnpacker)
// 	if err != nil {
// 		error.Handle(r.Context(), err, w, t.Container.Adapters.Log)
// 		return
// 	}

// 	// validate unpacked data
// 	errs := validator.Validate(sampleUnpacker)
// 	if errs != nil {
// 		error.HandleValidationErrors(r.Context(), errs, w, t.Container.Adapters.Log)
// 		return
// 	}

// 	// bind unpacked data to entities and pass to use case
// 	result, err := sample.Get()
// 	if err != nil {
// 		error.Handle(r.Context(), err, w, t.Container.Adapters.Log)
// 		return
// 	}

// 	// map result from use case to transformer
// 	var tr []transformers.TestTransformer
// 	for _, v := range result {
// 		tr = append(tr, transformers.TestTransformer{
// 			ID:   v.ID,
// 			Name: v.Name,
// 		})
// 	}

// 	// send response
// 	response.Send(w, response.Transform(tr), http.StatusOK)
// }
