package config

// AppConfig holds application configurations.
type AppConfig struct {
	Name     string       `yaml:"name"`
	Mode     string       `yaml:"mode"`
	Host     string       `yaml:"host"`
	Port     int          `yaml:"port"`
	Timezone string       `yaml:"timezone"`
	Metrics  MetricConfig `yaml:"metrics"`
}
