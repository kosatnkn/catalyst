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
)

func main() {
	// parse all configurations
	settings := config.Settings{
		Dir:    ".",
		Prefix: "CATALYST",
		Defaults: map[string]any{
			"rest.port": 8000,
		},
	}
	c, err := config.Parse(infra.Config{}, settings)
	if err != nil {
		msgNPanic(nil, errors.Join(errors.New("main: config parse failed"), err).Error())
	}
	cfg := c.(infra.Config)

	// resolve the container using parsed configurations
	ctr, err := infra.NewResolvedContainer(cfg)
	if err != nil {
		msgNPanic(nil, errors.Join(errors.New("main: error resolving container"), err).Error())
	}
	defer func() {
		ctr.Logger.Info(context.Background(), "destroying container")
		if err := ctr.Destroy(); err != nil {
			msgNPanic(ctr.Logger, errors.Join(errors.New("main: error destroying container"), err).Error())
		}
	}()

	// start service
	splash := strings.Join(append([]string{metadata.BaseInfo()}, metadata.BuildInfo()...), ", ")
	ctr.Logger.Info(context.Background(), splash)

	// start the REST server to handle requests
	restSrv, err := rest.NewServer(cfg.Rest, ctr)
	if err != nil {
		msgNPanic(ctr.Logger, errors.Join(errors.New("main: error creating REST server"), err).Error())
	}
	ctr.Logger.Info(context.Background(), "starting REST server")
	if err := restSrv.Start(); err != nil {
		msgNPanic(ctr.Logger, errors.Join(errors.New("main: error starting REST server"), err).Error())
	}
	defer func() {
		ctr.Logger.Info(context.Background(), "stopping REST server")
		if err := restSrv.Stop(); err != nil {
			msgNPanic(ctr.Logger, errors.Join(errors.New("main: error stopping REST server"), err).Error())
		}
	}()

	// enable graceful shutdown
	sig := make(chan os.Signal, 1)
	// accept graceful shutdowns when quit via SIGINT (Ctrl+C) or SIGTERM (Ctrl+/)
	// SIGKILL or SIGQUIT will not be caught
	signal.Notify(sig, syscall.SIGINT, syscall.SIGTERM)

	// block until a registered signal is received
	<-sig
	ctr.Logger.Info(context.Background(), "stopping service")
	// NOTE: We have used a service shutdown pattern leveraging `defer`.
	// This way, whenever a critical resource is created we do the destruction of that resource
	// right after the creation with a `defer`. This way when the `main` function returns or panics
	// it will destroy all attached resources properly.
}

// msgNPanic a convenient function to perform a logging and a panic to reduce redundancy.
func msgNPanic(l log.Logger, msg string) {
	msg = infra.FormatMsg(msg)

	if l == nil {
		fmt.Println(msg)
	} else {
		l.Error(context.Background(), msg)
	}

	panic("main")
}
