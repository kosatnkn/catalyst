package config

import (
	"fmt"

	"gopkg.in/yaml.v2"
)

// Parse parses all configuration to a single Config object.
func Parse() *Config {

	return &Config{
		AppConfig:      parseAppConfig(),
		DBConfig:       parseDBConfig(),
		LogConfig:      parseLogConfig(),
		ServicesConfig: parseServicesConfig(),
	}
}

// parseAppConfig parses application configurations.
func parseAppConfig() AppConfig {

	content := read("app.yaml")

	cfg := AppConfig{}

	err := yaml.Unmarshal(content, &cfg)

	if err != nil {
		panic(fmt.Sprintf("error: %v", err))
	}

	return cfg
}

// parseLogConfig parses logger configurations.
func parseLogConfig() LogConfig {

	content := read("logger.yaml")

	cfg := LogConfig{}

	err := yaml.Unmarshal(content, &cfg)

	if err != nil {
		panic(fmt.Sprintf("error: %v", err))
	}

	return cfg
}

// parseDBConfig parses database configurations.
func parseDBConfig() DBConfig {

	content := read("database.yaml")

	cfg := DBConfig{}

	err := yaml.Unmarshal(content, &cfg)

	if err != nil {
		panic(fmt.Sprintf("error: %v", err))
	}

	return cfg
}

// parseServicesConfig parses configurations of all services
func parseServicesConfig() []ServiceConfig {

	content := read("services.yaml")

	cfgs := []ServiceConfig{}

	err := yaml.Unmarshal(content, &cfgs)

	if err != nil {
		panic(fmt.Sprintf("error: %v", err))
	}

	return cfgs
}
