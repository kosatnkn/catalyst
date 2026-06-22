package profile

// Server defines the lifecycle interface for the telemetry server.
type Server interface {
	// Start begins listening on the configured port in a background goroutine.
	// It returns an error immediately if the listener cannot be bound (e.g. port
	// already in use), so the caller can fail fast at startup.
	Start() error

	// Stop performs a graceful shutdown, waiting for in-flight pprof requests
	// (such as long CPU profiles) to complete before returning.
	Stop() error
}
