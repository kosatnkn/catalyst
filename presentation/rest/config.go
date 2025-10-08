package rest

import (
	"context"
	"fmt"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/kosatnkn/catalyst/v3/infra"
)

func configure(cfg infra.RESTConfig, ctr *infra.Container) {
	// set running mode
	if cfg.Release {
		gin.SetMode(gin.ReleaseMode)
	}

	// set logging format of debug messages
	gin.DebugPrintFunc = func(format string, values ...any) {
		if len(format) > 0 {
			ctr.Logger.Debug(context.Background(), format)
		}
		if len(values) > 0 {
			var msgs []string
			for msg := range values {
				msgs = append(msgs, fmt.Sprintf("%v", msg))
			}
			ctr.Logger.Debug(context.Background(), strings.Join(msgs, ","))
		}
	}

	// set logging format of debug route messages
	gin.DebugPrintRouteFunc = func(httpMethod, absolutePath, handlerName string, nuHandlers int) {
		ctr.Logger.Debug(context.Background(), fmt.Sprintf("method: %s, path: %s, handler: %s, handlerCount: %d",
			httpMethod, absolutePath, handlerName, nuHandlers))
	}
}
