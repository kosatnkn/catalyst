package transformers

// APITransformer is used to transform the response payload for API details.
type APITransformer struct {
	Name    string `json:"name"`
	Version string `json:"version"`
	Purpose string `json:"purpose"`
}
