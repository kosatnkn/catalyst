package config

// ServiceConfig holds service configurations.
type ServiceConfig struct {
	Name    string `yaml:"name"`
	URL     string `yaml:"url"`
	Timeout int    `yaml:"timeout"`
}
