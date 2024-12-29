package postgres

// Config contains common database configurations for all database connections.
type Config struct {
	Host     string `yaml:"host"`
	Port     int    `yaml:"port"`
	Database string `yaml:"database"`
	User     string `yaml:"user"`
	Password string `yaml:"password"`
	PoolSize int    `yaml:"pool_size"`
	Check    bool   `yaml:"check"`
}
