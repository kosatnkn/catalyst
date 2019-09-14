package transformers

// ErrorTransformer is used to transform the response payload for errors.
type ErrorTransformer struct {
	Type  string `json:"type,omitempty"`
	Code  int    `json:"code,omitempty"`
	Msg   string `json:"message"`
	Trace string `json:"trace,omitempty"`
}
