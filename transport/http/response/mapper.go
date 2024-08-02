package response

import (
	"github.com/kosatnkn/catalyst/v2/transport/http/response/mappers"
	"github.com/kosatnkn/catalyst/v2/transport/http/response/transformers"
)

// mapData wraps payload in a standard response payload object.
func mapData(data []interface{}) (m mappers.Payload) {
	// map to fields using data types
	for _, v := range data {
		switch v.(type) {
		case transformers.PaginatorTransformer:
			m.Paginator = v
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
