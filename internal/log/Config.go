package log

// Config holds application log configurations.
type Config struct {
	// Level defines the log level that determines whether to log a message or not.
	//
	// There are four log levels to choose from with associated granularity.
	//  ERROR: log messages with log level 'ERROR'
	//  WARN: log messages with log level 'ERROR' and 'WARN'
	//  DEBUG: log messages with log level 'ERROR', 'WARN' and 'DEBUG'
	//  INFO: log messages with log level 'ERROR', 'WARN', 'DEBUG' and 'INFO'
	Level string `yaml:"level" mapstructure:"level"`
}
