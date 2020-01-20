package response

import (
	"encoding/json"

	"github.com/kosatnkn/catalyst/app/transport/response/mappers"
)

// Transform transforms a dataset in to a relevent structure and marshal to JSON.
func Transform(data interface{}, t TransformerInterface, isCollection bool) []byte {

	tData := transformByCriteria(data, t, isCollection)

	message, _ := json.Marshal(wrapInDataMapper(tData))

	return message
}

// transformByCriteria transforms data either as an object or as a collection
// depending on the `isCollection` boolean value
func transformByCriteria(data interface{}, t TransformerInterface, isCollection bool) interface{} {

	if isCollection {
		return t.TransformAsCollection(data)
	}

	return t.TransformAsObject(data)
}

// wrapInDataMapper wraps payload in a data object.
func wrapInDataMapper(data interface{}) mappers.Data {

	wrapper := mappers.Data{}
	wrapper.Payload = data

	return wrapper
}
