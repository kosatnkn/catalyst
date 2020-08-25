package config

// MetricConfig holds application metric configurations.
type MetricConfig struct {
	Enabled bool   `yaml:"enabled"`
	Port    int    `yaml:"port"`
	Route   string `yaml:"route"`
}
