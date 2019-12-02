package container

import (
	"fmt"

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

	db, err := adapters.NewPostgresAdapter(cfg)
	if err != nil {
		panic(fmt.Sprintf("error: %v", err))
	}

	resolvedAdapters.DB = db
}

// resolveDBTransactionAdapter resolves the database transaction adapter.
func resolveDBTransactionAdapter() {

	tx := adapters.NewPostgresTxAdapter(resolvedAdapters.DB)

	resolvedAdapters.DBTrans = tx
}

// resolveLogAdapter resolves the logging adapter.
func resolveLogAdapter(cfg config.LogConfig) {

	la, err := adapters.NewLogAdapter(cfg)
	if err != nil {
		panic(fmt.Sprintf("error: %v", err))
	}

	resolvedAdapters.Log = la
}

// resolveValidatorAdapter resolves the validator adapter.
func resolveValidatorAdapter() {

	va, err := adapters.NewValidatorAdapter()
	if err != nil {
		panic(fmt.Sprintf("error: %v", err))
	}

	resolvedAdapters.Validator = va
}