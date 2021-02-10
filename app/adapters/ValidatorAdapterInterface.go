package adapters

import "github.com/kosatnkn/validator"

// ValidatorAdapterInterface is implemented by all validator adapters.
type ValidatorAdapterInterface interface {
	validator.AdapterInterface
}
