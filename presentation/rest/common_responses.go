package rest

import (
	"github.com/gin-gonic/gin"
	"github.com/kosatnkn/catalyst/v3/infra"
)

// responseError defines the structure of the error response payload.
func responseError(err error) gin.H {
	return gin.H{
		"error": infra.FormatMsg(err.Error()),
	}
}

// responseData defines the structure of the data response payload.
func responseData(data, paging any) gin.H {
	if paging == nil {
		return gin.H{
			"data": data,
		}
	}

	return gin.H{
		"data":   data,
		"paging": paging,
	}
}
