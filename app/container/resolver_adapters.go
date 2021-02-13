package container

import (
	"fmt"

	"github.com/kosatnkn/catalyst/app/adapters"
	"github.com/kosatnkn/catalyst/app/config"
	"github.com/kosatnkn/db"
	"github.com/kosatnkn/db/mysql"
	"github.com/kosatnkn/log"
	"github.com/kosatnkn/validator"
)

// resolveAdapters resolves all adapters.
func resolveAdapters(cfg *config.Config) Adapters {

	ats := Adapters{}

	ats.DB = resolveDBAdapter(cfg.DB)
	ats.DBTx = resolveDBTransactionAdapter(ats.DB)
	ats.Log = resolveLogAdapter(cfg.Log)
	ats.Validator = resolveValidatorAdapter()

	return ats
}

// resolveDBAdapter resolves the database adapter.
func resolveDBAdapter(cfg mysql.Config) adapters.DBAdapterInterface {

	db, err := mysql.NewAdapter(cfg)
	if err != nil {
		panic(fmt.Sprintf("error: %v", err))
	}

	return db
}

// resolveDBTransactionAdapter resolves the database transaction adapter.
func resolveDBTransactionAdapter(d adapters.DBAdapterInterface) adapters.DBTxAdapterInterface {

	return db.NewTxAdapter(d)
}

// resolveLogAdapter resolves the logging adapter.
func resolveLogAdapter(cfg log.Config) adapters.LogAdapterInterface {

	la, err := log.NewAdapter(cfg)
	if err != nil {
		panic(fmt.Sprintf("error: %v", err))
	}

	return la
}

// resolveValidatorAdapter resolves the validation adapter.
func resolveValidatorAdapter() adapters.ValidatorAdapterInterface {

	v, err := validator.NewAdapter()
	if err != nil {
		panic(fmt.Sprintf("error: %v", err))
	}

	return v
}
