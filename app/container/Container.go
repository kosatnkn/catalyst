package container

import (
	"github.com/kosatnkn/catalyst/v3/app/adapters"
	"github.com/kosatnkn/catalyst/v3/domain/boundary/repositories"
)

// Container holds all resolved dependencies that needs to be injected at run time.
type Container struct {
	Adapters     Adapters
	Repositories Repositories
}

// Adapters hold resolved adapter instances.
//
// These are wrappers around third party libraries. All adapters will be of a corresponding adapter interface type.
type Adapters struct {
	DB        adapters.DBAdapterInterface
	Log       adapters.LogAdapterInterface
	Validator adapters.ValidatorAdapterInterface
}

// Repositories hold resolved repository instances.
//
// These are used to connect to databases. All repositories will be of a corresponding repository interface type.
// They act as an abstraction layer between the application and the database.
// Generally a single repository represents a single table in the database along with all the actions
// that can be performed on that table.
type Repositories struct {
	SampleRepository repositories.SampleRepositoryInterface
}
