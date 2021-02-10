package container

import (
	"fmt"

	"github.com/kosatnkn/catalyst/app/config"
	"github.com/kosatnkn/db/mysql"
	"github.com/kosatnkn/log"
	"github.com/kosatnkn/validator"
)

var ra Adapters

// resolveAdapters resolves all adapters.
func resolveAdapters(cfg *config.Config) Adapters {

	resolveDBAdapter(cfg.DB)
	resolveDBTransactionAdapter()
	resolveLogAdapter(cfg.Log)
	resolveValidatorAdapter()

	return ra
}

// resolveDBAdapter resolves the database adapter.
func resolveDBAdapter(cfg mysql.Config) {

	db, err := mysql.NewAdapter(cfg)
	if err != nil {
		panic(fmt.Sprintf("error: %v", err))
	}

	ra.DB = db
}

// resolveDBTransactionAdapter resolves the database transaction adapter.
func resolveDBTransactionAdapter() {

	ra.DBTx = mysql.NewTxAdapter(ra.DB)
}

// resolveLogAdapter resolves the logging adapter.
func resolveLogAdapter(cfg log.Config) {

	la, err := log.NewAdapter(cfg)
	if err != nil {
		panic(fmt.Sprintf("error: %v", err))
	}

	ra.Log = la
}

// resolveValidatorAdapter resolves the validation adapter.
func resolveValidatorAdapter() {

	v, err := validator.NewAdapter()
	if err != nil {
		panic(fmt.Sprintf("error: %v", err))
	}

	ra.Validator = v
}
