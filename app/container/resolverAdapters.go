package container

import (
	"github.com/kosatnkn/catalyst/app/config"
	"github.com/kosatnkn/catalyst/externals/adapters"
)

var resolvedAdapters Adapters

// Resolve all adapters.
func resolveAdapters(cfg *config.Config) Adapters {

	resolveDBAdapter(cfg.DBConfig)
	resolveLogAdapter(cfg.LogConfig)

	return resolvedAdapters
}

// Resolve the database adapter.
func resolveDBAdapter(cfg config.DBConfig) {

	// pg := adapters.PostgresAdapter{}
	db, _ := adapters.NewPostgresAdapter(cfg)
	// db, _ := pg.New(cfg)

	resolvedAdapters.DB = db
}

func resolveLogAdapter(cfg config.LogConfig) {

	l := adapters.LogAdapter{}
	la, _ := l.New(cfg)

	resolvedAdapters.Log = la
}
