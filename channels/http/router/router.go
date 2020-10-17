package router

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"

	"github.com/kosatnkn/catalyst/app/container"
	"github.com/kosatnkn/catalyst/channels/http/controllers"
	"github.com/kosatnkn/catalyst/channels/http/middleware"
	"github.com/kosatnkn/catalyst/channels/http/response"
)

// Init initializes the router.
func Init(ctr *container.Container) *mux.Router {

	// create new router
	r := mux.NewRouter()

	// initialize middleware
	requestCheckerMiddleware := middleware.NewRequestCheckerMiddleware(ctr)
	requestAlterMidleware := middleware.NewRequestAlterMiddleware()
	metricsMidleware := middleware.NewMetricsMiddleware()

	// add middleware to router
	// NOTE: middleware will execute in the order they are added to the router

	// add metrics middleware first
	r.Use(metricsMidleware.Middleware)

	// add CORS middleware
	r.Use(mux.CORSMethodMiddleware(r))

	r.Use(requestCheckerMiddleware.Middleware)
	r.Use(requestAlterMidleware.Middleware)

	// initialize controllers
	sampleController := controllers.NewSampleController(ctr)

	// bind controller functions to routes

	// api info
	r.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {

		name := "Catalyst"
		version := "v2.1.0"
		purpose := "REST API base written in Golang"

		response.Send(
			w,
			[]byte(fmt.Sprintf(`{
				"name": "%s",
				"version": "%s",
				"purpose": "%s"
			  }`, name, version, purpose)),
			http.StatusOK)
	}).Methods(http.MethodGet)

	// sample
	r.HandleFunc("/samples", sampleController.Get).Methods(http.MethodGet)
	r.HandleFunc("/samples/{id:[0-9]+}", sampleController.GetByID).Methods(http.MethodGet)
	r.HandleFunc("/samples", sampleController.Add).Methods(http.MethodPost)
	r.HandleFunc("/samples/{id:[0-9]+}", sampleController.Edit).Methods(http.MethodPut)
	r.HandleFunc("/samples/{id:[0-9]+}", sampleController.Delete).Methods(http.MethodDelete)

	return r
}
