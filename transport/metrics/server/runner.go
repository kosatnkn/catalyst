package metrics

import (
	"context"
	"fmt"
	"net/http"
	"strconv"
	"time"

	"github.com/prometheus/client_golang/prometheus/promhttp"

	"github.com/kosatnkn/catalyst/v3/app/config"
	"github.com/kosatnkn/catalyst/v3/app/container"
	"github.com/kosatnkn/catalyst/v3/app/metrics"
)

// Run runs a server to exposes metrics as a separate Prometheus metric server.
func Run(cfg config.AppConfig, ctr *container.Container) *http.Server {
	if !cfg.Metrics.Enabled {
		return nil
	}

	// register defined metrics
	metrics.Register()

	mux := http.NewServeMux()
	mux.Handle(cfg.Metrics.Route, promhttp.Handler())

	srv := &http.Server{
		Addr: cfg.Host + ":" + strconv.Itoa(cfg.Metrics.Port),

		// good practice to set timeouts to avoid Slowloris attacks
		WriteTimeout: time.Second * 15,
		ReadTimeout:  time.Second * 15,
		IdleTimeout:  time.Second * 60,

		Handler: mux,
	}

	// run the server in a goroutine so that it doesn't block
	go func() {
		err := srv.ListenAndServe()
		if err != nil {
			fmt.Println(err)
		}
	}()

	fmt.Printf("Metrics server exposing metrics on %s%s ...\n", srv.Addr, cfg.Metrics.Route)

	return srv
}

// Stop stops the server.
func Stop(srv *http.Server) {
	if srv == nil {
		return
	}

	fmt.Println("Metrics server shutting down...")
	srv.Shutdown(context.Background())
}
