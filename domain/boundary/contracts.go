package boundary

import "context"

// Readiness is used to report readiness state of components of this layer
// to the infrastructure layer.
type Readiness interface {
	// SetReadiness sets the readiness state of the component.
	SetReadiness(component string, ready bool)
}

// DatabaseTx is the interface that any database transaction adapter attaching to the service should implement.
type DatabaseTx interface {
	// Identity returns an identifier for the adapter.
	Identity() string

	// WrapInTx runs the content of the function in a single transaction.
	WrapInTx(ctx context.Context, fn func(ctx context.Context) (any, error)) (any, error)

	// IsReadinessFail check for common readiness failure scenarios.
	IsReadinessFail(err error) bool
}
