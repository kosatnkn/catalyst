package main

import (
	"github.com/kosatnkn/catalyst/app/config"
	"github.com/kosatnkn/catalyst/app/container"
	"github.com/kosatnkn/catalyst/app/http/server"
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
	server.Run(cfg.AppConfig, ctr)
}
