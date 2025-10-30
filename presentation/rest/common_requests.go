package rest

import (
	"encoding/json"
	"errors"
	"fmt"
	"strings"

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
//	<url>?filters=["filter1:val","filter2:1","filter3:true"]
//	<url>?filters=["filter4:[1,2,3]","filter5:[\"a\",\"b\",\"c\"]","filter6:[1.2,2.4,4.5]"]
//	<url>?filters=["filter7:{\"key1\":\"val1\",\"key2\":\"val2\",\"key3\":[1,2,3]}"]
func filters(ctx *gin.Context) (map[string]any, error) {
	// default filters
	f := map[string]any{}

	filters := ctx.Query("filters")
	if filters == "" {
		return f, nil
	}

	var arr []string
	if err := json.Unmarshal([]byte(filters), &arr); err != nil {
		return f, errors.Join(errors.New("rest: error unmarshaling filters from query param"), err)
	}

	for _, item := range arr {
		parts := strings.SplitN(item, ":", 2) // only split into 2 parts
		if len(parts) != 2 {
			return f, fmt.Errorf("rest: filter '%s' is of invalid format", item)
		}

		key := strings.TrimSpace(parts[0])
		val := strings.TrimSpace(parts[1])

		f[key] = val
	}

	return f, nil
}

// accountRequest maps account details sent in the request payload.
type accountRequest struct {
	Owner    string  `json:"owner" binding:"required"`
	Currency string  `json:"currency"`
	Balance  float32 `json:"balance"`
}
