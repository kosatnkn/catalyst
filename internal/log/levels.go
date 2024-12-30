package log

import "github.com/rs/zerolog"

const (
	levelError string = "ERROR"
	levelWarn  string = "WARN"
	levelDebug string = "DEBUG"
	levelInfo  string = "INFO"
)

var granularity map[string]zerolog.Level = map[string]zerolog.Level{
	levelInfo:  zerolog.InfoLevel,
	levelDebug: zerolog.DebugLevel,
	levelWarn:  zerolog.WarnLevel,
	levelError: zerolog.ErrorLevel,
}
