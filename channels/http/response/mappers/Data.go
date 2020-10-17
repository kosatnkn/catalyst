package mappers

// Data is the base mapper for data payloads.
type Data struct {
	Payload interface{} `json:"data"`
}
