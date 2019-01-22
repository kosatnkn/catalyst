package router

import (
	"net/http"

	"github.com/gorilla/mux"

	"github.com/kosatnkn/catalyst/app/container"
	"github.com/kosatnkn/catalyst/app/controllers"
	"github.com/kosatnkn/catalyst/app/middleware"
)

// Init initializes the router.
func Init(container *container.Container) *mux.Router {

	// create new router
	r := mux.NewRouter()

	// initialize middleware
	requestCheckerMiddleware := middleware.RequestCheckerMiddleware{}
	requestCheckerMiddleware.Init(container)

	requestAlterMidleware := middleware.RequestAlterMiddleware{}

	metricsMidleware := middleware.MetricsMiddleware{}

	// add middleware to router
	// NOTE: middleware will execute in the order they are added to the router
	// add metrics middleware first
	r.Use(metricsMidleware.Middleware)
	r.Use(requestCheckerMiddleware.Middleware)
	r.Use(requestAlterMidleware.Middleware)

	// initialize controllers
	testController := controllers.TestController{
		Container: container,
	}

	// bind controller functions to routes
	r.HandleFunc("/", testController.TestFunc).Methods(http.MethodPost)

	return r
}
