package transformers

// ErrorTransformer is used to transform the response payload for errors.
type ErrorTransformer struct {
	Type string `json:"type,omitempty"`
	Code string `json:"code,omitempty"`
	Msg  string `json:"message"`
}
