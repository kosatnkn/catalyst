package container

import (
	"github.com/kosatnkn/catalyst/app/config"
	"github.com/kosatnkn/catalyst/externals/adapters"
)

var resolvedAdapters Adapters

// resolveAdapters resolves all adapters.
func resolveAdapters(cfg *config.Config) Adapters {

	resolveDBAdapter(cfg.DBConfig)
	resolveLogAdapter(cfg.LogConfig)

	return resolvedAdapters
}

// resolveDBAdapter resolves the database adapter.
func resolveDBAdapter(cfg config.DBConfig) {

	db, _ := adapters.NewPostgresAdapter(cfg)

	resolvedAdapters.DB = db
}

// resolveLogAdapter resolves the logging adapter.
func resolveLogAdapter(cfg config.LogConfig) {

	la, _ := adapters.NewLogAdapter(cfg)

	resolvedAdapters.Log = la
}
