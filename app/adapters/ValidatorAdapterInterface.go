package adapters

// ValidatorAdapterInterface is implemented by all validator adapters.
type ValidatorAdapterInterface interface {
	// Validate validates fields of a struct.
	Validate(data any) map[string]string

	// ValidateField validates a single variable.
	ValidateField(name string, value any, rules string) map[string]string
}
