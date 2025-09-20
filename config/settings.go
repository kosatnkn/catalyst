package config

// Settings contain additional configuration values used by the config parser in order to parse the config structure
// that is provided.
// Under the hood the config parser uses 'github.com/spf13/viper' to parse configurations.
// So the config struct 'MUST' contain `mapstructure` tags in order for parser to do the mapping.
type Settings struct {
	// Directory in which the configuration file is located in.
	Dir string

	// Prefix used for the environment variable.
	// This will prevent accidental use of environment variables that were
	// not intended to be used for the service.
	// ex: `SERVICENAME_USER` make more sense than using `USER` because it
	// prevents you from using a system variable that you have not defined yourself.
	Prefix string

	// Default values to be used in case they are not provided.
	// Keys of the `Settings.Default` map should be dot (.) concatenated.
	// For a config struct like this,
	//
	//	type Config struct {
	//		App AppConfig `mapstructure:"app"`
	//	}
	//
	//	type AppConfig struct {
	//	 Name     string       `mapstructure:"name"`
	//		Metrics  MetricConfig `mapstructure:"metrics"`
	//	}
	//
	//	type MetricConfig struct {
	//		Enabled bool   `mapstructure:"enabled"`
	//		Port    int    `mapstructure:"port"`
	//		Route   string `mapstructure:"route"`
	//	}
	//
	// Use `app.metrics.enabled` as the key to set a default value for the metrics enabled configuration.
	Defaults map[string]any
}
