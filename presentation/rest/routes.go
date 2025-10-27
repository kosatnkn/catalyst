package rest

import (
	"github.com/gin-gonic/gin"
	"github.com/kosatnkn/catalyst/v3/infra"
)

func newRouter(cfg infra.RESTConfig, ctr *infra.Container) *gin.Engine {
	configure(cfg, ctr)

	router := gin.New()
	// middleware
	router.Use(gin.Recovery(),
		loggerMiddleware(ctr),
	)

	// api info handlers
	registerInfoHandlers(router)
	// usecase handlers
	registerAccountHandlers(router, ctr)

	return router
}

func registerInfoHandlers(router *gin.Engine) {
	router.GET("/", infoHandler)
	router.GET("/healthz", healthHandler)
	router.GET("/ready", readyHandler)
}

func registerAccountHandlers(router *gin.Engine, ctr *infra.Container) {
	accounts := newAccountController(ctr)

	a := router.Group("/accounts")
	a.GET("", accounts.Get)
	a.POST("", accounts.Create)
}
