package container

import (
	"fmt"

	"github.com/kosatnkn/catalyst/app/config"
	"github.com/kosatnkn/catalyst/externals/adapters"
)

var ra Adapters

// resolveAdapters resolves all adapters.
func resolveAdapters(cfg *config.Config) Adapters {

	resolveDBAdapter(cfg.DBConfig)
	resolveDBTransactionAdapter()
	resolveLogAdapter(cfg.LogConfig)
	resolveValidatorAdapter()

	return ra
}

// resolveDBAdapter resolves the database adapter.
func resolveDBAdapter(cfg config.DBConfig) {

	db, err := adapters.NewMySQLAdapter(cfg)
	if err != nil {
		panic(fmt.Sprintf("error: %v", err))
	}

	ra.DB = db
}

// resolveDBTransactionAdapter resolves the database transaction adapter.
func resolveDBTransactionAdapter() {

	ra.DBTx = adapters.NewMySQLTxAdapter(ra.DB)
}

// resolveLogAdapter resolves the logging adapter.
func resolveLogAdapter(cfg config.LogConfig) {

	la, err := adapters.NewLogAdapter(cfg)
	if err != nil {
		panic(fmt.Sprintf("error: %v", err))
	}

	ra.Log = la
}

// resolveValidatorAdapter resolves the validation adapter.
func resolveValidatorAdapter() {

	v, err := adapters.NewValidatorAdapter()
	if err != nil {
		panic(fmt.Sprintf("error: %v", err))
	}

	ra.Validator = v
}
