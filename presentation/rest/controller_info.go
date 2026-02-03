package rest

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/kosatnkn/catalyst/v3/infra"
	"github.com/kosatnkn/catalyst/v3/metadata"
)

// infoController is the collection of handlers used to retrieve information about the service.
type infoController struct {
	ctr *infra.Container
}

// newInfoController creates a new instance of the controller.
func newInfoController(ctr *infra.Container) *infoController {
	return &infoController{
		ctr: ctr,
	}
}

// Info returns details of the microservice.
func (c *infoController) Info(ctx *gin.Context) {
	data := gin.H{
		"name":    metadata.Name(),
		"version": metadata.BuildInfo(),
		"purpose": metadata.BaseInfo(),
	}

	ctx.JSON(http.StatusOK, responseData(data, nil))
}

// Health to be used by the Kubernetes liveliness probe.
func (c *infoController) Health(ctx *gin.Context) {
	ctx.String(http.StatusOK, "ok")
}

// Ready to be used by the Kubernetes readiness probe.
func (c *infoController) Ready(ctx *gin.Context) {
	if c.ctr.Lifecycle.Ready() {
		ctx.String(http.StatusOK, "ready")
		return
	}

	ctx.String(http.StatusServiceUnavailable, "not ready")
}
