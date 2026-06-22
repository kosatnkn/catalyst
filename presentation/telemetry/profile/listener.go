package profile

import (
	"net"
)

// newListener creates a bound TCP listener for addr.
// Binding eagerly (before Serve) means a port-conflict error surfaces in
// Start() and propagates to the caller, rather than being swallowed inside a
// goroutine where it can only be logged.
func newListener(addr string) (net.Listener, error) {
	return net.Listen("tcp", addr)
}
