package infra

import (
	"context"
	"fmt"
	"maps"
	"strings"
	"sync"
	"time"
)

// componentCheckerFn checks for the readiness of a component.
type componentCheckerFn func() (bool, error)

// Readiness contains the current state of different
// components of the service.
type Readiness struct {
	logger   Logger
	mu       sync.RWMutex
	states   map[string]bool
	checkers map[string]componentCheckerFn
}

// newReadiness creates a new instance.
func newReadiness(l Logger) *Readiness {
	return &Readiness{
		logger:   l,
		states:   make(map[string]bool),
		checkers: make(map[string]componentCheckerFn),
	}
}

// SetComponent updates readiness for a component.
func (r *Readiness) SetComponent(name string, ready bool) {
	r.mu.Lock()
	prev := r.states[name]
	r.states[name] = ready
	checker := r.checkers[name]
	r.mu.Unlock()

	// only trigger recovery when ready state changes from true to false
	if prev && !ready && checker != nil {
		ctx := context.Background()
		r.logger.Warn(ctx, fmt.Sprintf("Component '%s' is not ready, checking for readiness", name))
		go r.recoverComponent(ctx, name, checker)
	}
}

// Ready returns true only if all components are ready.
func (r *Readiness) Ready() bool {
	r.mu.RLock()
	defer r.mu.RUnlock()

	if len(r.states) == 0 {
		return false
	}

	for _, ready := range r.states {
		if !ready {
			return false
		}
	}

	return true
}

// RegisterComponentChecker registers a componentCheckerFn for the component with name.
func (r *Readiness) RegisterComponentChecker(name string, checker componentCheckerFn) {
	r.mu.Lock()
	r.checkers[name] = checker
	r.states[name] = true
	r.mu.Unlock()

	r.SetComponent(name, false)
}

// Snapshot returns current component states.
func (r *Readiness) Snapshot() map[string]bool {
	r.mu.RLock()
	defer r.mu.RUnlock()

	// NOTE: since maps are reference types returning `r.state`
	// will give access to it from the outside.
	// cloning prevents this by returning a copy of the map
	return maps.Clone(r.states)
}

// String returns current component states as a string.
func (r *Readiness) String() string {
	r.mu.RLock()
	defer r.mu.RUnlock()

	parts := make([]string, 0, len(r.states))
	for k, v := range r.states {
		parts = append(parts, fmt.Sprintf("%s: %t", k, v))
	}

	return strings.Join(parts, ", ")
}

// recoverComponent tries to recover the component state.
func (r *Readiness) recoverComponent(ctx context.Context, name string, checker componentCheckerFn) {
	ticker := time.NewTicker(5 * time.Second)
	defer ticker.Stop()

	for {
		select {
		case <-ctx.Done():
			r.logger.Warn(ctx, fmt.Sprintf("Cancelled readiness for component '%s'", name))
			return

		case <-ticker.C:
			ready, err := checker()
			if err != nil || !ready {
				r.logger.Warn(ctx, fmt.Sprintf("Readiness check fails for component '%s' due to error '%s'", name, err))
				continue
			}

			r.SetComponent(name, true)
			r.logger.Info(ctx, fmt.Sprintf("Readiness check passed for component '%s'", name))

			return
		}
	}
}
