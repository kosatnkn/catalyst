package postgres

// Config contains configuration parameters for Postgres.
type Config struct {
	Host     string `yaml:"host" mapstructure:"host"`
	Port     int    `yaml:"port" mapstructure:"port"`
	Database string `yaml:"database" mapstructure:"database"`
	User     string `yaml:"user" mapstructure:"user"`
	Password string `yaml:"password" mapstructure:"password"`
	PoolSize int    `yaml:"poolsize" mapstructure:"poolsize"`
	Check    bool   `yaml:"check" mapstructure:"check"`
}
