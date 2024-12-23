package server

import (
	"context"
	"fmt"
	"net/http"
	"strconv"
	"time"

	"github.com/kosatnkn/catalyst/v3/app/config"
	"github.com/kosatnkn/catalyst/v3/app/container"
	"github.com/kosatnkn/catalyst/v3/transport/http/router"
)

// Run runs the http server.
func Run(cfg config.AppConfig, ctr *container.Container) *http.Server {
	// initialize the router
	r := router.Init(ctr)

	srv := &http.Server{
		Addr: cfg.Host + ":" + strconv.Itoa(cfg.Port),

		// good practice to set timeouts to avoid Slowloris attacks
		WriteTimeout: time.Second * 15,
		ReadTimeout:  time.Second * 15,
		IdleTimeout:  time.Second * 60,

		// pass the instance of gorilla/mux in
		Handler: r,
	}

	// run the server in a goroutine so that it doesn't block
	go func() {
		err := srv.ListenAndServe()
		if err != nil {
			fmt.Println(err)
		}
	}()

	fmt.Printf("HTTP server listening on %v ...\n", srv.Addr)

	return srv
}

// Stop stops the server.
func Stop(srv *http.Server) {
	fmt.Println("HTTP server shutting down...")
	srv.Shutdown(context.Background())
}
