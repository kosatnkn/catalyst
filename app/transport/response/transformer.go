package response

import (
	"encoding/json"
)

// Transform transforms a dataset in to a relevent structure and marshal to JSON.
func Transform(data interface{}, t TransformerInterface, isCollection bool) []byte {

	var tData interface{}

	if isCollection {
		tData = t.TransformAsCollection(data)
	}

	tData = t.TransformAsObject(data)

	message, _ := json.Marshal(wrapInDataMapper(tData))

	return message
}

// wrapInDataMapper wraps payload in a data object.
func wrapInDataMapper(data interface{}) Data {

	wrapper := Data{}
	wrapper.Payload = data

	return wrapper
}
