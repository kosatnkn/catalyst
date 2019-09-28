package config

// Config holds all other config structs.
type Config struct {
	AppConfig      AppConfig
	DBConfig       DBConfig
	LogConfig      LogConfig
	ServicesConfig []ServiceConfig
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

// LogConfig holds application log configurations.
type LogConfig struct {
	Level     string `yaml:"level"`
	Remote    bool   `yaml:"remote_log"`
	FilePath  bool   `yaml:"file_path_enabled"`
	Colors    bool   `yaml:"colors"`
	Console   bool   `yaml:"console"`
	File      bool   `yaml:"file"`
	Directory string `yaml:"directory"`
}

// DBConfig holds database configurations.
type DBConfig struct {
	Host     string `yaml:"host"`
	Port     int    `yaml:"port"`
	Database string `yaml:"database"`
	User     string `yaml:"user"`
	Password string `yaml:"password"`
	PoolSize int    `yaml:"pool_size"`
}

// ServiceConfig holds service configurations.
type ServiceConfig struct {
	Name    string `yaml:"name"`
	URL     string `yaml:"url"`
	Timeout int    `yaml:"timeout"`
}
