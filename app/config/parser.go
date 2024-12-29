package config

import (
	"fmt"
	"os"

	"github.com/kosatnkn/catalyst/v3/internal/db/postgres"
	"github.com/kosatnkn/catalyst/v3/internal/log"
	"gopkg.in/yaml.v2"
)

// Parse parses all configuration to a single Config object.
// `cfgDir` is the path of the configuration directory.
func Parse(cfgDir string) *Config {
	// set config directory
	dir := getConfigDir(cfgDir)

	return &Config{
		App: parseAppConfig(dir),
		DB:  parseDBConfig(dir),
		Log: parseLogConfig(dir),
	}
}

// parseAppConfig parses application configurations.
func parseAppConfig(dir string) AppConfig {
	cfg := AppConfig{}
	parseConfig(dir+"app.yaml", &cfg)

	return cfg
}

// parseLogConfig parses logger configurations.
func parseLogConfig(dir string) log.Config {
	cfg := log.Config{}
	parseConfig(dir+"logger.yaml", &cfg)

	return cfg
}

// parseDBConfig parses database configurations.
func parseDBConfig(dir string) postgres.Config {
	cfg := postgres.Config{}
	parseConfig(dir+"database.yaml", &cfg)

	return cfg
}

// parseConfig reads configuration values from the given file and
// populates the given config struct.
func parseConfig(file string, unpacker interface{}) {
	content := read(file)

	err := yaml.Unmarshal(content, unpacker)
	if err != nil {
		panic(fmt.Sprintf("error: %v", err))
	}
}

// getConfigDir returns config directory path after analyzing and correcting.
func getConfigDir(dir string) string {
	// get last char of dir path
	c := dir[len(dir)-1]
	if os.IsPathSeparator(c) {
		return dir
	}

	return dir + string(os.PathSeparator)
}
