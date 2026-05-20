package infra

// Readiness is the interface that any readiness updater attaching to the service should implement.
type Readiness interface {
	// SetReadiness updates readiness for a component.
	SetReadiness(component string, ready bool)

	// Ready returns true only if all components are ready.
	Ready() bool

	// RegisterCheckerFn registers a callback function for the component.
	RegisterCheckerFn(component string, checker func() (bool, error))

	// Snapshot returns current component states.
	Snapshot() map[string]bool

	// String returns current component states as a string.
	String() string
}
