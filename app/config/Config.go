package config

import (
	"github.com/kosatnkn/catalyst/v3/internal/db/postgres"
	"github.com/kosatnkn/catalyst/v3/internal/log"
)

// Config is the master config struct that holds all other config structs.
type Config struct {
	App AppConfig
	DB  postgres.Config
	Log log.Config
}

// AppConfig holds application configurations.
type AppConfig struct {
	Name     string       `yaml:"name"`
	Mode     string       `yaml:"mode"`
	Host     string       `yaml:"host"`
	Port     int          `yaml:"port"`
	Timezone string       `yaml:"timezone"`
	Metrics  MetricConfig `yaml:"metrics"`
}

// MetricConfig holds application metric configurations.
type MetricConfig struct {
	Enabled bool   `yaml:"enabled"`
	Port    int    `yaml:"port"`
	Route   string `yaml:"route"`
}
