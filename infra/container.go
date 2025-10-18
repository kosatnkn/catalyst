package infra

import (
	"errors"

	"github.com/kosatnkn/catalyst-pkgs/persistence"
	"github.com/kosatnkn/catalyst-pkgs/persistence/postgres"
	"github.com/kosatnkn/catalyst-pkgs/telemetry/log"
	"github.com/kosatnkn/catalyst-pkgs/telemetry/log/loggerjson"
	"github.com/kosatnkn/catalyst/v3/domain/boundary"
	external "github.com/kosatnkn/catalyst/v3/persistence/postgres"
)

// Container is a simple implementation of a dependency inversion container.
//
// There is nothing fancy in this implementation. Dependencies are resolved manually.
// This can be replaced with a DI container of your choice.
type Container struct {
	dbAdapter        persistence.DatabaseAdapter
	Logger           log.Logger
	AccountRetriever boundary.AccountRetriever
	AccountPersister boundary.AccountPersister
}

// NewResolvedContainer returns a fresh instance of the container after resolving dependencies.
func NewResolvedContainer(cfg Config) (*Container, error) {
	c := &Container{}

	logger, err := loggerjson.NewLoggerJSON(cfg.Log)
	if err != nil {
		return nil, errors.Join(errors.New("container: error creating logger"), err)
	}
	c.Logger = logger

	pg, err := postgres.NewDatabaseAdapterPostgres(cfg.Database)
	if err != nil {
		return nil, errors.Join(errors.New("container: error creating postgres adapter"), err)
	}
	c.dbAdapter = pg

	c.AccountRetriever = external.NewAccountRetrieverPostgres(c.dbAdapter)
	c.AccountPersister = external.NewAccountPersisterPostgres(c.dbAdapter)

	return c, nil
}

// Destroy will execute the ordered destruction of objects in the container.
//
// NOTE: For this container destruct is done manually.
func (c *Container) Destroy() error {
	err := c.dbAdapter.Destruct()
	if err != nil {
		return errors.Join(errors.New("container: error destroying database adapter"), err)
	}

	return nil
}
