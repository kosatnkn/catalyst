package main

import (
	"fmt"
	"os"
	"os/signal"

	"github.com/kosatnkn/catalyst/v3/app/config"
	"github.com/kosatnkn/catalyst/v3/app/container"
	"github.com/kosatnkn/catalyst/v3/app/splash"
	httpServer "github.com/kosatnkn/catalyst/v3/transport/http/server"
	metricsServer "github.com/kosatnkn/catalyst/v3/transport/metrics/server"
)

func main() {
	// parse all configurations
	cfg := config.Parse("./configs")

	// show splash screen when starting
	splash.Show(splash.StyleDefault, cfg)

	// resolve the container using parsed configurations
	ctr := container.Resolve(cfg)

	fmt.Println("Starting...")
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
	fmt.Println("\nStopping...")
	httpServer.Stop(hsrv)
	metricsServer.Stop(msrv)

	// release resources
	ctr.Destruct()

	fmt.Println("Done")

	os.Exit(0)
}
