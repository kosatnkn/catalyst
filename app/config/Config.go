package config

import (
	"github.com/kosatnkn/catalyst/v3/internal/db/postgres"
	"github.com/kosatnkn/catalyst/v3/internal/log"
)

// Config is the master config struct that holds all other config structs.
type Config struct {
	App AppConfig       `mapstructure:"app"`
	DB  postgres.Config `mapstructure:"db"`
	Log log.Config      `mapstructure:"log"`
}

// AppConfig holds application configurations.
type AppConfig struct {
	Name     string       `mapstructure:"name"`
	Mode     string       `mapstructure:"mode"`
	Host     string       `mapstructure:"host"`
	Port     int          `mapstructure:"port"`
	Timezone string       `mapstructure:"timezone"`
	Metrics  MetricConfig `mapstructure:"metrics"`
}

// MetricConfig holds application metric configurations.
type MetricConfig struct {
	Enabled bool   `mapstructure:"enabled"`
	Port    int    `mapstructure:"port"`
	Route   string `mapstructure:"route"`
}
