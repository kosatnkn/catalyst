package config

import (
	"fmt"
	"os"

	"gopkg.in/yaml.v2"
)

var configDir string

// Parse parses all configuration to a single Config object.
func Parse(cfgDir string) *Config {

	setConfigDir(cfgDir)

	return &Config{
		AppConfig:      parseAppConfig(),
		DBConfig:       parseDBConfig(),
		LogConfig:      parseLogConfig(),
		ServicesConfig: parseServicesConfig(),
	}
}

// setConfigDir sets the configuration directory.
func setConfigDir(dir string) {

	// get last char of dir path
	c := dir[len(dir)-1]

	if os.IsPathSeparator(c) {

		configDir = dir

		return
	}

	configDir = dir + string(os.PathSeparator)
}

// parseAppConfig parses application configurations.
func parseAppConfig() AppConfig {

	cfg := AppConfig{}

	parseConfig("app.yaml", &cfg)

	return cfg
}

// parseLogConfig parses logger configurations.
func parseLogConfig() LogConfig {

	cfg := LogConfig{}

	parseConfig("logger.yaml", &cfg)

	return cfg
}

// parseDBConfig parses database configurations.
func parseDBConfig() DBConfig {

	cfg := DBConfig{}

	parseConfig("database.yaml", &cfg)

	return cfg
}

// parseServicesConfig parses configurations of all services
func parseServicesConfig() []ServiceConfig {

	cfgs := []ServiceConfig{}

	parseConfig("services.yaml", &cfgs)

	return cfgs
}

// parseConfig reads configuration values from the given file and
// populates the given config struct.
func parseConfig(fileName string, unpacker interface{}) {

	content := read(fileName)

	err := yaml.Unmarshal(content, unpacker)
	if err != nil {
		panic(fmt.Sprintf("error: %v", err))
	}
}
