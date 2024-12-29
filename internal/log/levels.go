package log

const (
	levelError string = "ERROR"
	levelWarn  string = "WARN"
	levelDebug string = "DEBUG"
	levelInfo  string = "INFO"
)

var granularity map[string]int = map[string]int{
	levelInfo:  1,
	levelDebug: 2,
	levelWarn:  3,
	levelError: 4,
}
