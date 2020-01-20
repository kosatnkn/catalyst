package main

import (
	"context"
	"os"
	"os/signal"
	"time"

	"github.com/kosatnkn/catalyst/app/config"
	"github.com/kosatnkn/catalyst/app/container"
	"github.com/kosatnkn/catalyst/app/http/server"
	"github.com/kosatnkn/catalyst/app/metrics"
	"github.com/kosatnkn/catalyst/app/splash"
)

func main() {

	// show splash screen when starting
	splash.Show(splash.StyleDefault)

	// parse all configurations
	cfg := config.Parse("./config")

	// resolve the container using parsed configurations
	ctr := container.Resolve(cfg)

	// start the server to handle requests
	srv := server.Run(cfg.AppConfig, ctr)

	// expose application metrics
	metrics.Expose(cfg.AppConfig, ctr)

	// enable graceful shutdown
	c := make(chan os.Signal, 1)

	// accept graceful shutdowns when quit via SIGINT (Ctrl+C)
	// SIGKILL, SIGQUIT or SIGTERM (Ctrl+/) will not be caught
	signal.Notify(c, os.Interrupt)

	// block until a registered signal is received
	<-c

	// create a deadline to wait for
	var wait time.Duration

	ctx, cancel := context.WithTimeout(context.Background(), wait)
	defer cancel()

	// gracefully stop the server
	server.Stop(ctx, srv, ctr)

	os.Exit(0)
}
