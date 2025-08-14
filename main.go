package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"

	"github.com/kosatnkn/catalyst/v3/app/config"
	"github.com/kosatnkn/catalyst/v3/app/container"
	"github.com/kosatnkn/catalyst/v3/app/splash"
	httpServer "github.com/kosatnkn/catalyst/v3/app/transport/http/server"
	metricsServer "github.com/kosatnkn/catalyst/v3/app/transport/metrics/server"
)

func main() {
	// show splash prompt when starting
	splash.Show()

	// parse all configurations
	cfg, err := config.Parse(".")
	if err != nil {
		fmt.Println("Error parsing configurations:", err)
		os.Exit(1)
	}

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
	// accept graceful shutdowns when quit via SIGINT (Ctrl+C) or SIGTERM (Ctrl+/)
	// SIGKILL or SIGQUIT will not be caught
	signal.Notify(c, syscall.SIGINT, syscall.SIGTERM)

	// block until a registered signal is received
	<-c

	// Shutdown in the reverse order of initialization.
	fmt.Println("\nStopping...")
	httpServer.Stop(hsrv)
	metricsServer.Stop(msrv)

	// release resources
	fmt.Println("Releasing resources...")
	ctr.Destruct()

	fmt.Println("Done")

	os.Exit(0)
}
