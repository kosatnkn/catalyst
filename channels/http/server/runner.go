package server

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"strconv"
	"time"

	"github.com/kosatnkn/catalyst/v2/app/config"
	"github.com/kosatnkn/catalyst/v2/app/container"
	"github.com/kosatnkn/catalyst/v2/channels/http/router"
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

		// pass our instance of gorilla/mux in
		Handler: r,
	}

	// run our server in a goroutine so that it doesn't block
	go func() {

		err := srv.ListenAndServe()
		if err != nil {
			log.Println(err)
			panic("Service shutting down unexpectedly...")
		}
	}()

	fmt.Println("Service started...")
	fmt.Println(fmt.Sprintf("Listening on %v ...", srv.Addr))

	return srv
}

// Stop stops the server.
func Stop(ctx context.Context, srv *http.Server) {

	fmt.Println("Service shutting down...")

	srv.Shutdown(ctx)
}
