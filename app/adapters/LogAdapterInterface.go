package adapters

import "context"

// LogAdapterInterface is implemented by all logging adapters.
type LogAdapterInterface interface {
	// AddTraceID attaches a trace id to context that can be later read by the logger.
	AddTraceID(ctx context.Context, id string) context.Context

	// AppendTracePoint appends the given trace point to a trace path in context that can be later read by the logger.
	AppendTracePoint(ctx context.Context, point string) context.Context

	// Error logs a message as of error type.
	Error(ctx context.Context, message string, options ...interface{})

	// Debug logs a message as of debug type.
	Debug(ctx context.Context, message string, options ...interface{})

	// Info logs a message as of information type.
	Info(ctx context.Context, message string, options ...interface{})

	// Warn logs a message as of warning type.
	Warn(ctx context.Context, message string, options ...interface{})

	// Destruct will close the logger gracefully releasing all resources.
	Destruct()
}
