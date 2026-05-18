package postgres

import (
	"context"
	"database/sql"
	"errors"
	"net"
	"strings"

	"github.com/kosatnkn/catalyst/v3/persistence"
)

type readinessCheckHelper struct {
	ready persistence.DatabaseReadinessAdapter
}

// newReadinessCheckHelper creates a new instance.
func newReadinessCheckHelper(ready persistence.DatabaseReadinessAdapter) *readinessCheckHelper {
	return &readinessCheckHelper{
		ready: ready,
	}
}

// withReadinessCheck pipe the error through
// readiness check logic before returning it.
func (h *readinessCheckHelper) withReadinessCheck(err error) error {
	if err != nil && h.isReadinessFail(err) {
		h.ready.SetComponent(Identity, false)
	}

	return err
}

// isReadinessFail check for common readiness failure scenarios
// for Postgres database connections.
func (h *readinessCheckHelper) isReadinessFail(err error) bool {
	// Context timeout usually indicates infra/connectivity issues.
	if errors.Is(err, context.DeadlineExceeded) {
		return true
	}
	// database/sql connection became unusable.
	if errors.Is(err, sql.ErrConnDone) {
		return true
	}
	// Generic network-level failures.
	if _, ok := errors.AsType[net.Error](err); ok {
		return true
	}
	// Fallback string matching for lower-level driver/network errors.
	msg := strings.ToLower(err.Error())
	transientIndicators := []string{
		"connection refused",
		"connection reset",
		"broken pipe",
		"unexpected eof",
		"eof",
		"no such host",
		"server closed the connection",
		"network is unreachable",
		"i/o timeout",
		"dial tcp",
		"driver: bad connection",
	}
	for _, s := range transientIndicators {
		if strings.Contains(msg, s) {
			return true
		}
	}

	return false
}
