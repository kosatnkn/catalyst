package main

import (
	"github.com/kosatnkn/catalyst/app/config"
	"github.com/kosatnkn/catalyst/app/container"
	"github.com/kosatnkn/catalyst/app/router"
	"github.com/kosatnkn/catalyst/app/splash"
	"github.com/kosatnkn/catalyst/server"
)

func main() {

	// show splash screen when starting
	splash.Show(splash.StyleShadow)

	// parse all configurations
	cfg := config.Parse("./config")

	// resolve the container using parsed configurations
	ctr := container.Resolve(cfg)

	// initialize the router
	r := router.Init(ctr)

	// start the server to handle requests
	server.Run(cfg.AppConfig, r, ctr)
}
