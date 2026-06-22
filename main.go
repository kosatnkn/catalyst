package main

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"strings"
	"syscall"

	"github.com/kosatnkn/catalyst-pkgs/config"
	"github.com/kosatnkn/catalyst-pkgs/telemetry/log"
	"github.com/kosatnkn/catalyst/v3/infra"
	"github.com/kosatnkn/catalyst/v3/metadata"
	"github.com/kosatnkn/catalyst/v3/presentation/rest"
	"github.com/kosatnkn/catalyst/v3/presentation/telemetry/profile"
)

func main() {
	// create a single context to be used by all logging calls
	lctx := context.Background()

	// parse all configurations
	settings := config.Settings{
		Dir:    ".",
		Prefix: "CATALYST",
		Defaults: map[string]any{
			"app.name":     metadata.Name(),
			"app.mode":     "DEBUG",
			"app.timezone": "UTC",
			"rest.port":    8000,
			"rest.wait":    "5s",
			"rest.release": false,
			"log.level":    "INFO",
			// NOTE: The profiling server is OPTIONAL
			// if profiler is added, add sensible default
			// values here so that the service
			// can be started with minimal config
			"telemetry.profiler.enabled":     false,
			"telemetry.profiler.port":        6060,
			"telemetry.profiler.readtimeout": 60,
		},
	}
	c, err := config.Parse(infra.Config{}, settings)
	if err != nil {
		msgNPanic(lctx, nil, fmt.Errorf("main: config parse failed, %w", err).Error())
	}
	cfg := c.(infra.Config)

	// resolve the container using parsed configurations
	ctr, err := infra.NewResolvedContainer(cfg)
	if err != nil {
		msgNPanic(lctx, nil, fmt.Errorf("main: error resolving container, %w", err).Error())
	}
	defer func() {
		ctr.Logger.Info(lctx, "destroying container")
		if err := ctr.Destroy(); err != nil {
			msgNPanic(lctx, ctr.Logger, fmt.Errorf("main: error destroying container, %w", err).Error())
		}
	}()

	// start service
	splash := strings.Join(append([]string{metadata.BaseInfo()}, metadata.BuildInfo()...), ", ")
	ctr.Logger.Info(lctx, splash)

	// start the REST server to handle requests
	restSrv, err := rest.NewServer(cfg.Rest, ctr)
	if err != nil {
		msgNPanic(lctx, ctr.Logger, fmt.Errorf("main: error creating REST server, %w", err).Error())
	}
	ctr.Logger.Info(lctx, "starting REST server")
	if err := restSrv.Start(); err != nil {
		msgNPanic(lctx, ctr.Logger, fmt.Errorf("main: error starting REST server, %w", err).Error())
	}
	defer func() {
		ctr.Logger.Info(lctx, "stopping REST server")
		if err := restSrv.Stop(); err != nil {
			msgNPanic(lctx, ctr.Logger, fmt.Errorf("main: error stopping REST server, %w", err).Error())
		}
	}()

	// start profiling server
	// NOTE: The profiling server is OPTIONAL
	// If you don't need it remove this section.
	profSrv, err := profile.NewServer(cfg.Telemetry.Profile, ctr)
	if err != nil {
		msgNPanic(lctx, ctr.Logger, fmt.Errorf("main: error creating profiling server, %w", err).Error())
	}
	if err := profSrv.Start(); err != nil {
		msgNPanic(lctx, ctr.Logger, fmt.Errorf("main: error starting profiling server, %w", err).Error())
	}
	defer func() {
		if err := profSrv.Stop(); err != nil {
			msgNPanic(lctx, ctr.Logger, fmt.Errorf("main: error stopping profiling server, %w", err).Error())
		}
	}()

	// mark service as ready to accept connections
	ctr.Logger.Info(lctx, fmt.Sprintf("ready to accept connections on port %d", cfg.Rest.Port))

	// enable graceful shutdown
	sig := make(chan os.Signal, 1)
	// accept graceful shutdowns when quit via SIGINT (Ctrl+C) or SIGTERM (Ctrl+/)
	// SIGKILL or SIGQUIT will not be caught
	signal.Notify(sig, syscall.SIGINT, syscall.SIGTERM)
	defer signal.Stop(sig)

	// block until a registered signal is received
	<-sig
	ctr.Logger.Info(lctx, "stopping service")
	// NOTE: We have used a service shutdown pattern leveraging `defer`.
	// This way, whenever a critical resource is created we do the destruction of that resource
	// right after the creation with a `defer`. So when the `main` function returns or panics
	// it will destroy all attached resources properly.
}

// msgNPanic a convenient function to perform a logging and a panic to reduce redundancy.
func msgNPanic(ctx context.Context, l log.Logger, msg string) {
	msg = infra.FormatMsg(msg)

	if l == nil {
		fmt.Println(msg)
	} else {
		l.Error(ctx, msg)
	}

	panic("main")
}
