package response

import (
	"github.com/kosatnkn/catalyst/v3/app/transport/http/response/transformer"
)

// Transform transforms data either as an object or as a collection depending on the `isCollection` boolean value.
func Transform(data interface{}, t transformer.TransformerInterface, isCollection bool) (interface{}, error) {
	if isCollection {
		return t.TransformAsCollection(data)
	}

	return t.TransformAsObject(data)
}
