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

	ra.DBAdapter = db
}

// resolveDBTransactionAdapter resolves the database transaction adapter.
func resolveDBTransactionAdapter() {

	ra.DBTxAdapter = adapters.NewMySQLTxAdapter(ra.DBAdapter)
}

// resolveLogAdapter resolves the logging adapter.
func resolveLogAdapter(cfg config.LogConfig) {

	la, err := adapters.NewLogAdapter(cfg)
	if err != nil {
		panic(fmt.Sprintf("error: %v", err))
	}

	ra.LogAdapter = la
}

// resolveValidatorAdapter resolves the validation adapter.
func resolveValidatorAdapter() {

	v, err := adapters.NewValidatorAdapter()
	if err != nil {
		panic(fmt.Sprintf("error: %v", err))
	}

	ra.ValidatorAdapter = v
}
