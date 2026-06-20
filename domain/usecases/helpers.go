package usecases

import "github.com/kosatnkn/catalyst/v3/domain/boundary"

// withDBReadinessCheck pipe the error through database readiness check logic before returning it.
// This will set the database readiness to 'false' if the readiness check fails.
func withDBReadinessCheck(tx boundary.DatabaseTx, r boundary.Readiness, err error) error {
	if err != nil && tx.IsReadinessFail(err) {
		r.SetReadiness(tx.Identity(), false)
	}

	return err
}
