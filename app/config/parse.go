package config

import (
	"fmt"
	"os"

	"github.com/spf13/viper"
)

// Parse parses all configuration to a single Config object.
// `cfgDir` is the path of the configuration directory.
func Parse(cfgDir string) (*Config, error) {
	var config Config

	// set config directory
	dir := getConfigDir(cfgDir)

	// initialize viper
	viper.SetConfigName("config") // Name of the config file (without extension)
	viper.SetConfigType("yaml")   // Config file format
	viper.AddConfigPath(dir)      // Path to the directory containing the config file
	viper.SetEnvPrefix("CATALYST")
	viper.AutomaticEnv() // Enable environment variable override

	// set default values
	viper.SetDefault("app.name", "no-name")
	viper.SetDefault("app.mode", "DEBUG")
	viper.SetDefault("app.host", "")
	viper.SetDefault("app.port", "8080")
	viper.SetDefault("app.timezone", "UTC")
	viper.SetDefault("app.metrics.enabled", true)
	viper.SetDefault("app.metrics.port", 8081)
	viper.SetDefault("app.metrics.route", "/metrics")
	viper.SetDefault("db.host", "localhost")
	viper.SetDefault("db.port", 5432)
	viper.SetDefault("db.database", "postgres")
	viper.SetDefault("db.user", "postgres")
	viper.SetDefault("db.password", "")
	viper.SetDefault("db.pool_size", 10)
	viper.SetDefault("db.check", true)
	viper.SetDefault("log.level", "INFO")

	// try to read the config file
	if err := viper.ReadInConfig(); err != nil {
		fmt.Println("falling back to environment variables:", err)
	}

	// Unmarshal into the struct
	if err := viper.Unmarshal(&config); err != nil {
		return nil, fmt.Errorf("unable to decode into struct: %w", err)
	}

	fmt.Printf("Config: %+v\n", config)

	return &config, nil
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
