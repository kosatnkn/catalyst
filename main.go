package main

import (
	"context"
	"errors"
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
	if err := run(); err != nil {
		// run has already logged the error and all of its defers have completed
		os.Exit(1)
	}
}

// run starts the service and blocks until it is stopped.
// It returns a non nil error when the service failed (startup, runtime or teardown).
func run() (err error) {
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
		return logErr(lctx, nil, fmt.Errorf("main: config parse failed, %w", err))
	}
	cfg := c.(infra.Config)

	// resolve the container using parsed configurations
	ctr, err := infra.NewResolvedContainer(cfg)
	if err != nil {
		return logErr(lctx, nil, fmt.Errorf("main: error resolving container, %w", err))
	}
	defer func() {
		ctr.Logger.Info(lctx, "destroying container")
		if derr := ctr.Destroy(); derr != nil {
			err = errors.Join(err, logErr(lctx, ctr.Logger, fmt.Errorf("main: error destroying container, %w", derr)))
		}
	}()

	// start service
	splash := strings.Join(append([]string{metadata.BaseInfo()}, metadata.BuildInfo()...), ", ")
	ctr.Logger.Info(lctx, splash)

	// start the REST server to handle requests
	restSrv, err := rest.NewServer(cfg.Rest, ctr)
	if err != nil {
		return logErr(lctx, ctr.Logger, fmt.Errorf("main: error creating REST server, %w", err))
	}
	ctr.Logger.Info(lctx, "starting REST server")
	if err := restSrv.Start(); err != nil {
		return logErr(lctx, ctr.Logger, fmt.Errorf("main: error starting REST server, %w", err))
	}
	defer func() {
		ctr.Logger.Info(lctx, "stopping REST server")
		if serr := restSrv.Stop(); serr != nil {
			err = errors.Join(err, logErr(lctx, ctr.Logger, fmt.Errorf("main: error stopping REST server, %w", serr)))
		}
	}()

	// start profiling server
	// NOTE: The profiling server is OPTIONAL
	// If you don't need it remove this section.
	profSrv, err := profile.NewServer(cfg.Telemetry.Profile, ctr)
	if err != nil {
		return logErr(lctx, ctr.Logger, fmt.Errorf("main: error creating profiling server, %w", err))
	}
	if err := profSrv.Start(); err != nil {
		return logErr(lctx, ctr.Logger, fmt.Errorf("main: error starting profiling server, %w", err))
	}
	defer func() {
		if serr := profSrv.Stop(); serr != nil {
			err = errors.Join(err, logErr(lctx, ctr.Logger, fmt.Errorf("main: error stopping profiling server, %w", serr)))
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

	// block until a registered signal is received or a server fails
	// NOTE: a server failure here is unrecoverable (ex: the listener died),
	// request failures never reach this point.
	select {
	case <-sig:
		ctr.Logger.Info(lctx, "stopping service")
	case serr := <-restSrv.Err():
		return logErr(lctx, ctr.Logger, fmt.Errorf("main: REST server failed, %w", serr))
	case serr := <-profSrv.Err():
		return logErr(lctx, ctr.Logger, fmt.Errorf("main: profiling server failed, %w", serr))
	}

	// NOTE: We have used a service shutdown pattern leveraging `defer`.
	// Whenever a critical resource is created we do the destruction of that resource
	// right after the creation with a `defer`. So when `run` returns (normally or with an error)
	// it will destroy all attached resources properly, and `main` exits non zero on error.
	return nil
}

// logErr logs the error (through the logger, or stdout before the logger exists) and returns it.
func logErr(ctx context.Context, l log.Logger, err error) error {
	msg := infra.FormatMsg(err.Error())

	if l == nil {
		fmt.Println(msg)
	} else {
		l.Error(ctx, msg)
	}

	return err
}
