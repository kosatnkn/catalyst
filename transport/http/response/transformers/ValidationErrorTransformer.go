package transformers

// ValidationErrorTransformer is used to transform the response payload for validation errors.
type ValidationErrorTransformer struct {
	Type string `json:"type,omitempty"`
	Msg  any    `json:"message,omitempty"`
}
