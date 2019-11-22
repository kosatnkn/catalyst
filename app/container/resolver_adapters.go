package container

import (
	"github.com/kosatnkn/catalyst/app/config"
	"github.com/kosatnkn/catalyst/externals/adapters"
)

var resolvedAdapters Adapters

// resolveAdapters resolves all adapters.
func resolveAdapters(cfg *config.Config) Adapters {

	resolveDBAdapter(cfg.DBConfig)
	resolveDBTransactionAdapter()
	resolveLogAdapter(cfg.LogConfig)
	resolveValidatorAdapter()

	return resolvedAdapters
}

// resolveDBAdapter resolves the database adapter.
func resolveDBAdapter(cfg config.DBConfig) {

	db, _ := adapters.NewPostgresAdapter(cfg)

	resolvedAdapters.DB = db
}

// resolveDBTransactionAdapter resolves the database transaction adapter.
func resolveDBTransactionAdapter() {

	tx := adapters.NewPostgresTxAdapter(resolvedAdapters.DB)

	resolvedAdapters.DBTrans = tx
}

// resolveLogAdapter resolves the logging adapter.
func resolveLogAdapter(cfg config.LogConfig) {

	la, _ := adapters.NewLogAdapter(cfg)

	resolvedAdapters.Log = la
}

// resolveValidatorAdapter resolves the validator adapter.
func resolveValidatorAdapter() {

	va, _ := adapters.NewValidatorAdapter()

	resolvedAdapters.Validator = va
}
