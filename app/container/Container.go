package container

import (
	"github.com/kosatnkn/catalyst/domain/boundary/adapters"
	"github.com/kosatnkn/catalyst/domain/boundary/repositories"
	"github.com/kosatnkn/catalyst/domain/boundary/services"
)

// Container holds all resolved dependencies that needs to be injected at run time.
type Container struct {
	Adapters     Adapters
	Repositories Repositories
	Services     Services
}

// Adapters hold resolved adapter instances.
// These are wrappers around third party libraries. All adapters will be of a corresponding adapter interface type.
type Adapters struct {
	DB  adapters.DBAdapterInterface
	Log adapters.LogAdapterInterface
}

// Repositories hold resolved repository instances.
// These are used to connect to databases. All repositories will be of a corresponding repository interface type.
// They act as an abstraction layer between the application and the database.
// Generally a single repository represents a single table in the database along with all the actions
// that can be performed on that table.
type Repositories struct {
	SampleRepository repositories.SampleRepositoryInterface
}

// Services hold resolved service instances.
// These are abstractions to third party APIs. All services will be of a corresponding service interface type.
type Services struct {
	SampleService services.SampleServiceInterface
}
