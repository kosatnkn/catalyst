package controllers

import (
	"net/http"
	"strings"

	"github.com/kosatnkn/catalyst/v3/app/container"
	"github.com/kosatnkn/catalyst/v3/app/transport/http/controllers"
	"github.com/kosatnkn/catalyst/v3/metadata"
	"github.com/kosatnkn/catalyst/v3/transport/http/response/transformers"
)

// APIController contains controller logic for endpoints.
type APIController struct {
	*controllers.Controller
}

// NewAPIController creates a new instance of the controller.
func NewAPIController(c *container.Container) *APIController {
	return &APIController{
		Controller: controllers.NewController(c),
	}
}

// GetInfo return basic details of the API.
func (ctl *APIController) GetInfo(w http.ResponseWriter, r *http.Request) {
	// transform
	tr := transformers.APITransformer{
		Name:    "Catalyst",
		Version: strings.ReplaceAll(metadata.BuildInfo(), "\n", " "),
		Purpose: "A REST API base written in Golang",
	}

	// send response
	ctl.SendResponse(w, http.StatusOK, tr)
}
