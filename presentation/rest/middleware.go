package rest

import (
	"context"
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/kosatnkn/catalyst/v3/infra"
)

func loggerMiddleware(ctr *infra.Container) gin.HandlerFunc {
	return gin.LoggerWithFormatter(func(params gin.LogFormatterParams) string {
		if params.ErrorMessage != "" {
			ctr.Logger.Error(context.Background(),
				fmt.Sprintf("method: %s, status: %d, path: %s, err: %s",
					params.Method,
					params.StatusCode,
					params.Path,
					infra.FormatMsg(params.ErrorMessage),
				),
			)
			return ""
		}

		ctr.Logger.Debug(context.Background(),
			fmt.Sprintf("method: %s, status: %d, path: %s",
				params.Method,
				params.StatusCode,
				params.Path,
			),
		)
		return ""
	})
}
