package config

import (
	"github.com/kosatnkn/db/mysql"
)

// Config is the master config struct that holds all other config structs.
type Config struct {
	App      AppConfig
	DB       mysql.Config
	Log      LogConfig
	Services []ServiceConfig
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

// LogConfig holds application log configurations.
type LogConfig struct {
	Level     string `yaml:"level"`
	Colors    bool   `yaml:"colors"`
	Console   bool   `yaml:"console"`
	File      bool   `yaml:"file"`
	Directory string `yaml:"directory"`
}

// MetricConfig holds application metric configurations.
type MetricConfig struct {
	Enabled bool   `yaml:"enabled"`
	Port    int    `yaml:"port"`
	Route   string `yaml:"route"`
}

// ServiceConfig holds service configurations.
type ServiceConfig struct {
	Name    string `yaml:"name"`
	URL     string `yaml:"url"`
	Timeout int    `yaml:"timeout"`
}
