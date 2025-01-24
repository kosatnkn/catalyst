package router

import (
	"net/http"

	"github.com/gorilla/mux"

	"github.com/kosatnkn/catalyst/v3/app/container"
	"github.com/kosatnkn/catalyst/v3/app/transport/http/middleware"
	"github.com/kosatnkn/catalyst/v3/transport/http/controllers"
)

// TODO: rewrite this using standard library.

// Init initializes the router.
func Init(ctr *container.Container) *mux.Router {
	// create new router
	r := mux.NewRouter()

	// add middleware to the router
	//
	// NOTE: middleware will execute in the order they are added to the router
	// add metrics middleware first
	r.Use(middleware.NewMetricsMiddleware().Middleware)
	// enable CORS
	r.Use(mux.CORSMethodMiddleware(r))
	r.Use(middleware.NewCORSMiddleware().Middleware)
	// other middleware
	r.Use(middleware.NewRequestAlterMiddleware(ctr).Middleware)
	r.Use(middleware.NewRequestCheckerMiddleware(ctr).Middleware)

	// initialize controllers
	apiController := controllers.NewAPIController(ctr)
	sampleController := controllers.NewSampleController(ctr)

	// bind controller functions to routes
	// api info
	r.HandleFunc("/", apiController.GetInfo).Methods(http.MethodGet)
	// sample
	r.HandleFunc("/samples", sampleController.Get).Methods(http.MethodGet)
	r.HandleFunc("/samples/{id:[0-9]+}", sampleController.GetByID).Methods(http.MethodGet)
	r.HandleFunc("/samples", sampleController.Add).Methods(http.MethodPost)
	r.HandleFunc("/samples/{id:[0-9]+}", sampleController.Edit).Methods(http.MethodPut)
	r.HandleFunc("/samples/{id:[0-9]+}", sampleController.Delete).Methods(http.MethodDelete)

	return r
}
