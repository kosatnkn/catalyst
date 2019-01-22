package response

// Data is the base mapper for data payloads.
type Data struct {
	Payload interface{} `json:"data"`
}

// Error is the base mapper for error payloads.
type Error struct {
	Payload interface{} `json:"errors"`
}
