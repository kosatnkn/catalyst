package controllers

import (
	"net/http"

	"github.com/kosatnkn/catalyst/app/container"
	"github.com/kosatnkn/catalyst/app/error"
	"github.com/kosatnkn/catalyst/app/transport/request"
	"github.com/kosatnkn/catalyst/app/transport/request/unpackers"
	"github.com/kosatnkn/catalyst/app/transport/response"
	"github.com/kosatnkn/catalyst/app/transport/response/transformers"
	"github.com/kosatnkn/catalyst/app/validator"
	"github.com/kosatnkn/catalyst/domain/globals"
	"github.com/kosatnkn/catalyst/domain/usecases"
)

// TestController contains controller logic for endpoints.
type TestController struct {
	Container *container.Container
}

// TestFunc handles POST request to "/" endpoint.
func (t *TestController) TestFunc(w http.ResponseWriter, r *http.Request) {

	// append a prefix value to the context passed within the request
	r = r.WithContext(globals.AppendToContextPrefix(r.Context(), "TestController.TestFunc"))

	// unpack request
	userUnpacker := unpackers.UserUnpacker{}
	err := request.Unpack(r, &userUnpacker)
	if err != nil {
		error.Handle(r.Context(), err, w, t.Container.Adapters.Log)
		return
	}

	// validate unpacked data
	errs := validator.Validate(userUnpacker)
	if errs != nil {
		error.HandleValidationErrors(r.Context(), errs, w, t.Container.Adapters.Log)
		return
	}

	// initialize use case
	test := usecases.TestUseCase{TestRepository: t.Container.Repositories.TestRepository}

	// bind unpacked data to entities and pass to use case
	result, err := test.TestDB()
	if err != nil {
		error.Handle(r.Context(), err, w, t.Container.Adapters.Log)
		return
	}

	// map result from use case to transformer
	var tr []transformers.TestTransformer
	for _, v := range result {
		tr = append(tr, transformers.TestTransformer{
			ID:   v.ID,
			Name: v.Name,
		})
	}

	// send response
	response.Send(w, response.Transform(tr), http.StatusOK)
}
