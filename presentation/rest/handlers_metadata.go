package rest

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/kosatnkn/catalyst/v3/metadata"
)

// NOTE:

// infoHandler returns details of the microservice.
func infoHandler(c *gin.Context) {
	data := gin.H{
		"name":    metadata.Name(),
		"version": metadata.BuildInfo(),
		"purpose": metadata.BaseInfo(),
	}
	c.JSON(http.StatusOK, responseData(data, nil))
}

// healthHandler to be used by the Kubernetes liveliness probe.
// TODO: implement logic
func healthHandler(c *gin.Context) {
	c.String(http.StatusOK, "ok")
}

// readyHandler to be used by the Kubernetes readiness probe.
// TODO: implement logic
func readyHandler(c *gin.Context) {
	c.String(http.StatusOK, "ready")
}
