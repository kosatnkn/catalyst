package main

import (
	"context"
	"errors"
	"fmt"
	"os"
	"os/signal"
	"strings"
	"syscall"
	"time"

	"github.com/kosatnkn/catalyst-pkgs/config"
	"github.com/kosatnkn/catalyst-pkgs/telemetry/log"
	"github.com/kosatnkn/catalyst/v3/infra"
	"github.com/kosatnkn/catalyst/v3/metadata"
	"github.com/kosatnkn/catalyst/v3/presentation/rest"
	"github.com/kosatnkn/catalyst/v3/presentation/telemetry/profile"
)

// set settings for the configuration parser
var confSettings = config.Settings{
	Dir:    ".",
	Prefix: "CATALYST",
	Defaults: map[string]any{
		"app.name":     metadata.Name(),
		"app.mode":     "DEBUG",
		"app.timezone": "UTC",
		"app.wait":     "5s",
		"rest.port":    8000,
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

// main entrypoint.
func main() {
	if err := run(); err != nil {
		// run has already logged the error and all of its defers have completed
		os.Exit(1)
	}
}

// run starts the service and blocks until it is stopped.
// It returns a non nil error when the service failed (startup, runtime or teardown).
//
// NOTE: `retErr` is a named result so that the deferred teardown funcs can report errors.
// A value returned from a deferred func is discarded, but a deferred func can assign to a named result.
// On `return X`, Go first sets `retErr = X`, then runs the defers (LIFO), and only then hands
// `retErr` to the caller. So each teardown defer does `retErr = errors.Join(retErr, ...)`
// to add its error without losing the one that caused the return.
// It is intentionally named `retErr` (not `err`) so that it is never shadowed by a local `err`.
func run() (retErr error) {
	// create a single context to be used by all logging calls
	lCtx := context.Background()

	// catch SIGINT (Ctrl+C) and SIGTERM from the very start, so that a signal received
	// while starting up still goes through the deferred teardown
	// NOTE: SIGKILL cannot be caught
	sigCtx, stop := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM)
	defer stop()

	c, err := config.Parse(infra.Config{}, confSettings)
	if err != nil {
		return logErr(lCtx, nil, fmt.Errorf("main, run: config parse failed, %w", err))
	}
	cfg := c.(infra.Config)

	// parse here so that a bad value fails at startup instead of breaking the graceful shutdown
	wait, err := time.ParseDuration(cfg.App.Wait)
	if err != nil {
		return logErr(lCtx, nil, fmt.Errorf("main, run: invalid app wait duration %q, %w", cfg.App.Wait, err))
	}

	// a single deadline (`app.wait`) shared by every teardown step
	// NOTE: it is started lazily by the first teardown step that asks for it, because teardown can begin
	// either after a signal or from an early return midway through startup.
	// Teardown runs in this goroutine only, so no synchronization is needed.
	var stopCtx context.Context
	var stopCancel context.CancelFunc
	shutdownCtx := func() context.Context {
		if stopCtx == nil {
			stopCtx, stopCancel = context.WithTimeout(context.Background(), wait)
		}
		return stopCtx
	}
	// registered before any resource defer, so that it runs last
	defer func() {
		if stopCancel != nil {
			stopCancel()
		}
	}()

	// resolve the container using parsed configurations
	ctr, err := infra.NewResolvedContainer(cfg)
	if err != nil {
		return logErr(lCtx, nil, fmt.Errorf("main, run: error resolving container, %w", err))
	}
	defer func() {
		ctr.Logger.Info(lCtx, "destroying container")
		if err := ctr.Destroy(); err != nil {
			retErr = errors.Join(retErr, logErr(lCtx, ctr.Logger, fmt.Errorf("main, run: error destroying container, %w", err)))
		}
	}()

	// start service
	splash := strings.Join(append([]string{metadata.BaseInfo()}, metadata.BuildInfo()...), ", ")
	ctr.Logger.Info(lCtx, splash)

	// start the REST server to handle requests
	restSrv, err := rest.NewServer(cfg.Rest, ctr)
	if err != nil {
		return logErr(lCtx, ctr.Logger, fmt.Errorf("main, run: error creating REST server, %w", err))
	}
	if err := restSrv.Start(); err != nil {
		return logErr(lCtx, ctr.Logger, fmt.Errorf("main, run: error starting REST server, %w", err))
	}
	defer func() {
		if err := restSrv.Stop(shutdownCtx()); err != nil {
			retErr = errors.Join(retErr, logErr(lCtx, ctr.Logger, fmt.Errorf("main, run: error stopping REST server, %w", err)))
		}
	}()

	// start profiling server
	// NOTE: The profiling server is OPTIONAL
	// If you don't need it remove this section.
	profSrv, err := profile.NewServer(cfg.Telemetry.Profile, ctr)
	if err != nil {
		return logErr(lCtx, ctr.Logger, fmt.Errorf("main, run: error creating profiling server, %w", err))
	}
	if err := profSrv.Start(); err != nil {
		return logErr(lCtx, ctr.Logger, fmt.Errorf("main, run: error starting profiling server, %w", err))
	}
	defer func() {
		if err := profSrv.Stop(shutdownCtx()); err != nil {
			retErr = errors.Join(retErr, logErr(lCtx, ctr.Logger, fmt.Errorf("main, run: error stopping profiling server, %w", err)))
		}
	}()

	// mark service as ready to accept connections
	ctr.Logger.Info(lCtx, fmt.Sprintf("ready to accept connections on port %d", cfg.Rest.Port))

	// block until a registered signal is received or a server fails
	// NOTE: a server failure here is unrecoverable (ex: the listener died),
	// request failures never reach this point.
	select {
	case <-sigCtx.Done():
		// restore default signal behaviour so a second Ctrl+C force quits a hung shutdown
		stop()
		ctr.Logger.Info(lCtx, fmt.Sprintf("stopping service, %v", context.Cause(sigCtx)))
	case serr := <-restSrv.Err():
		return logErr(lCtx, ctr.Logger, fmt.Errorf("main, run: REST server failed, %w", serr))
	case serr := <-profSrv.Err():
		return logErr(lCtx, ctr.Logger, fmt.Errorf("main, run: profiling server failed, %w", serr))
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
