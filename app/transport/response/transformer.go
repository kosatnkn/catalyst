package response

import (
	"encoding/json"
)

// Transform transforms a dataset in to a relevent structure and marshal to JSON.
func Transform(data interface{}) []byte {

	wrapper := Data{}
	wrapper.Payload = data

	message, _ := json.Marshal(wrapper)

	return message
}
