package mappers

// Payload is the base mapper for data payloads.
type Payload struct {
	Data      any `json:"data"`
	Paginator any `json:"paginator,omitempty"`
}
