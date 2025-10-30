package rest

import (
	"encoding/json"
	"errors"
	"fmt"

	"github.com/gin-gonic/gin"
)

// paging extracts pagination data from the provided param and send it as a map.
//
//	This function expects 'paging' to be in the following formats.
//	<url>?paging={"page":1,"size":10}
//	<url>?paging={"page":1}  // 'size' defaults to 10
//	<url>?paging={"size":10} // 'page' defaults to 1
func paging(ctx *gin.Context) (map[string]uint32, error) {
	// default paging
	paging := map[string]uint32{
		"page": 1,
		"size": 10,
	}

	p := ctx.Query("paging")
	if p == "" {
		return paging, nil
	}

	var pgn map[string]uint32
	if err := json.Unmarshal([]byte(p), &pgn); err != nil {
		return paging, errors.Join(errors.New("rest: error unmarshaling 'paging' from query param"), err)
	}
	for k, v := range pgn {
		switch k {
		case "page", "size":
			paging[k] = v
		default:
			return paging, fmt.Errorf("rest: invalid key '%s' in paging, only accept 'page', 'size", k)
		}
	}

	return paging, nil
}

// filters extracts filter data from the provided param and send it as a map.
//
//	This function expects filters to be in the following format.
//	<url>?filters={"name":"exam","tags":[1,2,3]}
func filters(ctx *gin.Context) (map[string]any, error) {
	// default filters
	filters := map[string]any{}

	f := ctx.Query("filters")
	if f == "" {
		return filters, nil
	}

	if err := json.Unmarshal([]byte(f), &filters); err != nil {
		return filters, errors.Join(errors.New("rest: error unmarshaling filters from query param"), err)
	}

	return filters, nil
}

// accountRequest maps account details sent in the request payload.
type accountRequest struct {
	Owner    string  `json:"owner" binding:"required"`
	Currency string  `json:"currency"`
	Balance  float32 `json:"balance"`
}
