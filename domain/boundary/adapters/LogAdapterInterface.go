package adapters

import (
	"context"
)

// LogAdapterInterface is implemented by all logging adapters.
type LogAdapterInterface interface {

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
