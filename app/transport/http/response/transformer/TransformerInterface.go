package transformer

// TransformerInterface is the interface implemented by all transformers.
type TransformerInterface interface {
	// TransformAsObject map data to a transformer object.
	TransformAsObject(data any) (any, error)

	// TransformAsCollection map data to a collection of transformer objects.
	TransformAsCollection(data any) (any, error)
}
