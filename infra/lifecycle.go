package infra

import "sync/atomic"

// lifecycle contains the current state of the service.
type lifecycle struct {
	ready atomic.Bool
}

// newLifecycle creates a new instance.
func newLifecycle() *lifecycle {
	return &lifecycle{}
}

// SetReady marks the ready state as the passed in state.
func (l *lifecycle) SetReady(val bool) {
	l.ready.Store(val)
}

// Ready retrieves the current state.
func (l *lifecycle) Ready() bool {
	return l.ready.Load()
}
