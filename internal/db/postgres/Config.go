package postgres

// Config contains common database configurations for all database connections.
type Config struct {
	Host     string `yaml:"host" mapstructure:"host"`
	Port     int    `yaml:"port" mapstructure:"port"`
	Database string `yaml:"database" mapstructure:"database"`
	User     string `yaml:"user" mapstructure:"user"`
	Password string `yaml:"password" mapstructure:"password"`
	PoolSize int    `yaml:"pool_size" mapstructure:"pool_size"`
	Check    bool   `yaml:"check" mapstructure:"check"`
}
