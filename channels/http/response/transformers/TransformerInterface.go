package transformers

// TransformerInterface is the interface implemented by all transformers.
type TransformerInterface interface {

	// TransformAsObject map data to a transformer object.
	TransformAsObject(data interface{}) (interface{}, error)

	// TransformAsCollection map data to a collection of transformer objects.
	TransformAsCollection(data interface{}) (interface{}, error)
}
