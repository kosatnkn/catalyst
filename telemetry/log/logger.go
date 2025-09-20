package log

import "context"

// Context key type to be used with contexts.
type key string

// ID is the universally unique identifier key to be used with context.
const ID key = "id"

// TraceKey is the key to add an additional trace values to the context.
const TraceKey key = "trace"

const (
	LevelError string = "ERROR"
	LevelWarn  string = "WARN"
	LevelDebug string = "DEBUG"
	LevelInfo  string = "INFO"
)

// Logger is implemented by all logging adapters.
type Logger interface {
	// AddTraceID attaches a trace id to context that can be later read by the logger.
	AddTraceID(ctx context.Context, id string) context.Context

	// AppendTracePoint appends the given trace point to a trace path in context that can be later read by the logger.
	AppendTracePoint(ctx context.Context, point string) context.Context

	// Error logs a message as of error type.
	Error(ctx context.Context, message string)

	// Debug logs a message as of debug type.
	Debug(ctx context.Context, message string)

	// Info logs a message as of information type.
	Info(ctx context.Context, message string)

	// Warn logs a message as of warning type.
	Warn(ctx context.Context, message string)
}
