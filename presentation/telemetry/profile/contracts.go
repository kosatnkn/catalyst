package profile

import "context"

// Server defines the lifecycle interface for the telemetry server.
type Server interface {
	// Start begins listening on the configured port in a background goroutine.
	// It returns an error immediately if the listener cannot be bound (e.g. port
	// already in use), so the caller can fail fast at startup.
	Start() error

	// Stop performs a graceful shutdown, waiting for in-flight pprof requests
	// (such as long CPU profiles) until ctx is done, then closes remaining connections.
	Stop(ctx context.Context) error

	// Err returns a channel that receives an unrecoverable error if the server
	// stops serving after a successful Start. A nil channel (never fires) is
	// valid for servers that cannot fail at runtime.
	Err() <-chan error
}
