package server

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"strconv"
	"time"

	"github.com/kosatnkn/catalyst/app/config"
	"github.com/kosatnkn/catalyst/app/container"
	"github.com/kosatnkn/catalyst/app/http/router"
)

// Run runs the application server.
func Run(cfg config.AppConfig, ctr *container.Container) {

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

	// expose application metrics
	exposeMetrics(cfg, ctr)

	c := make(chan os.Signal, 1)

	// We'll accept graceful shutdowns when quit via SIGINT (Ctrl+C)
	// SIGKILL, SIGQUIT or SIGTERM (Ctrl+/) will not be caught.
	signal.Notify(c, os.Interrupt)

	// Block until we receive our signal
	<-c

	// Create a deadline to wait for
	var wait time.Duration

	ctx, cancel := context.WithTimeout(context.Background(), wait)
	defer cancel()

	// gracefully release all additional resources
	destruct(ctr)

	// Doesn't block if no connections, but will otherwise wait
	// until the timeout deadline.
	srv.Shutdown(ctx)
	fmt.Println("Service shutting down...")
	os.Exit(0)
}

// Gracefully close all additional resources.
func destruct(ctr *container.Container) {

	fmt.Println("Closing database connections...")
	ctr.Adapters.DBAdapter.Destruct()

	fmt.Println("Closing logger...")
	ctr.Adapters.LogAdapter.Destruct()
}
