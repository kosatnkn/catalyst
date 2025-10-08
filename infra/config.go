package infra

import (
	"github.com/kosatnkn/catalyst-pkgs/persistence/postgres"
	"github.com/kosatnkn/catalyst-pkgs/telemetry/log/loggerjson"
)

// Config is the master config struct that holds all other config structs.
type Config struct {
	App      AppConfig         `mapstructure:"app"`
	Rest     RESTConfig        `mapstructure:"rest"`
	Metrics  MetricConfig      `mapstructure:"metrics"`
	Log      loggerjson.Config `mapstructure:"log"`
	Database postgres.Config   `mapstructure:"db"`
}

// AppConfig holds application configurations.
type AppConfig struct {
	Name     string `mapstructure:"name"`
	Mode     string `mapstructure:"mode"`
	Timezone string `mapstructure:"timezone"`
}

type RESTConfig struct {
	// Port in which the REST server is running on.
	Port int `mapstructure:"port" validate:"min=80,max=65535"`
	// The time duration in which the server waits until shutdown when a terminate signal is received.
	// This is a string value that can be parsed in to a time duration by `time.ParseDuration` function,
	// i.e. "300ms", "1.5h", or "2h45m".
	// Valid time units are "ns", "us" (or "µs"), "ms", "s", "m", "h".
	Wait string `mapstructure:"wait"`
	// Denotes whether the server should run in release (production) mode or not.
	// The server will run in debug mode when this is set to false.
	Release bool `mapstructure:"release"`
}

// MetricConfig holds application metric configurations.
type MetricConfig struct {
	Enabled bool   `mapstructure:"enabled"`
	Port    int    `mapstructure:"port"`
	Route   string `mapstructure:"route"`
}
