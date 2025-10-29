package rest

import (
	"github.com/gin-gonic/gin"
)

// bindErrorToCtx is a helper function that sets the error
// in the gin context so that it is accessible by th logger
// and the response.
func bindErrorToCtx(ctx *gin.Context, err error, httpStatus int) {
	// read by the logger
	ctx.Error(err)

	// error response from rest server
	ctx.JSON(httpStatus, responseError(err))
}
