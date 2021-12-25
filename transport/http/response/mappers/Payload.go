package mappers

// Payload is the base mapper for data payloads.
type Payload struct {
	Data      interface{} `json:"data"`
	Paginator interface{} `json:"paginator,omitempty"`
}
