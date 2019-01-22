package adapters

import (
	"context"

	"github.com/kosatnkn/catalyst/app/config"
)

// LogdapterInterface is implemeted by all looging adapters.
type LogdapterInterface interface {

	// New creates a new instance of log adapter implementation.
	New(config config.LogConfig) (LogdapterInterface, error)

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
