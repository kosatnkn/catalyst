package infra

import (
	"errors"

	rd "github.com/kosatnkn/catalyst-pkgs/infra/readiness/basic"
	pg "github.com/kosatnkn/catalyst-pkgs/persistence/postgres"
	lg "github.com/kosatnkn/catalyst-pkgs/telemetry/log/loggerjson"
	"github.com/kosatnkn/catalyst/v3/domain/boundary"
	"github.com/kosatnkn/catalyst/v3/persistence"
	"github.com/kosatnkn/catalyst/v3/persistence/postgres"
)

// Container is a simple implementation of a dependency inversion container.
//
// There is nothing fancy in this implementation. Dependencies are resolved manually.
// This can be replaced with a DI container of your choice.
type Container struct {
	Logger           Logger
	Readiness        Readiness
	DBAdapter        persistence.DatabaseAdapter
	AccountRetriever boundary.AccountRetriever
	AccountPersister boundary.AccountPersister
}

// NewResolvedContainer returns a fresh instance of the container after resolving dependencies.
func NewResolvedContainer(cfg Config) (*Container, error) {
	var err error
	c := &Container{}

	// logger
	if c.Logger, err = lg.NewLoggerJSON(cfg.Log); err != nil {
		return nil, errors.Join(errors.New("container: error creating logger"), err)
	}

	// readiness
	// NOTE: Set up the readiness probe to enable service readiness querying.
	c.Readiness = rd.NewReadiness(c.Logger)

	// database
	if c.DBAdapter, err = pg.NewDatabaseAdapterPostgres(cfg.Database); err != nil {
		return nil, errors.Join(errors.New("container: error creating postgres adapter"), err)
	}
	c.Readiness.RegisterCheckerFn(
		c.DBAdapter.Identity(), // NOTE: reference the component identifier from one place
		func() (bool, error) { // NOTE: function to run in order to check readiness of component
			if err := c.DBAdapter.Ping(); err != nil {
				return false, err
			}
			return true, nil
		})

	// domain
	c.AccountRetriever = postgres.NewAccountRetrieverPostgres(c.DBAdapter, c.Readiness) // NOTE: pass in the readiness probe only in to objects that needs readiness tracking
	c.AccountPersister = postgres.NewAccountPersisterPostgres(c.DBAdapter)

	return c, nil
}

// Destroy will execute the ordered destruction of objects in the container.
//
// NOTE: For this container destruct is done manually.
func (c *Container) Destroy() error {
	err := c.DBAdapter.Destruct()
	if err != nil {
		return errors.Join(errors.New("container: error destroying database adapter"), err)
	}

	return nil
}
