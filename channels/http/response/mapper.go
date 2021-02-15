package response

import "github.com/kosatnkn/catalyst/channels/http/response/mappers"

// mapData wraps payload in a standard response payload object.
func mapData(data []interface{}) (m mappers.Payload) {

	for _, v := range data {

		switch v.(type) {
		default:
			m.Data = v
		}
	}

	return m
}

// mapErr wraps error in a standard error response object.
func mapErr(err interface{}) mappers.Error {

	return mappers.Error{
		Err: err,
	}
}
