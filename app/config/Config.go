package config

// Config is the master config struct that holds all other config structs.
type Config struct {
	AppConfig      AppConfig
	DBConfig       DBConfig
	LogConfig      LogConfig
	ServicesConfig []ServiceConfig
}
