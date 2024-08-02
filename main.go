package main

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"time"

	"github.com/kosatnkn/catalyst/v2/app/config"
	"github.com/kosatnkn/catalyst/v2/app/container"
	"github.com/kosatnkn/catalyst/v2/app/splash"
	httpServer "github.com/kosatnkn/catalyst/v2/transport/http/server"
	metricsServer "github.com/kosatnkn/catalyst/v2/transport/metrics/server"
)

func main() {
	// show splash screen when starting
	splash.Show(splash.StyleDefault)
	// parse all configurations
	cfg := config.Parse("./configs")
	// resolve the container using parsed configurations
	ctr := container.Resolve(cfg)

	fmt.Println("Service starting...")

	// start the server to handle http requests
	hsrv := httpServer.Run(cfg.App, ctr)
	// start the server to expose application metrics
	msrv := metricsServer.Run(cfg.App, ctr)

	fmt.Println("Ready")

	// enable graceful shutdown
	c := make(chan os.Signal, 1)
	// accept graceful shutdowns when quit via SIGINT (Ctrl+C)
	// SIGKILL, SIGQUIT or SIGTERM (Ctrl+/) will not be caught
	signal.Notify(c, os.Interrupt)

	// block until a registered signal is received
	<-c

	// Shutdown in the reverse order of initialization.
	fmt.Println("\nService stopping...")

	// create a deadline to wait for
	var wait time.Duration

	// Doesn't block if no connections, but will otherwise wait
	// until the timeout deadline.
	ctx, cancel := context.WithTimeout(context.Background(), wait)
	defer cancel()

	// gracefully stop the http server
	httpServer.Stop(ctx, hsrv)
	// gracefully stop metrics server
	metricsServer.Stop(ctx, msrv)
	// release resources
	ctr.Destruct()

	fmt.Println("Done")

	os.Exit(0)
}
