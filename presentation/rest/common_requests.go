package rest

import (
	"encoding/json"
	"errors"
	"fmt"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
)

// filters extracts filter data from the provided param and send it as a map.
//
//	This function expects filters to be in the following format.
//	<url>?filters=["filter1:val","filter2:1","filter3:true"]
//	<url>?filters=["filter4:[1,2,3]","filter5:[\"a\",\"b\",\"c\"]","filter6:[1.2,2.4,4.5]"]
//	<url>?filters=["filter7:{\"key1\":\"val1\",\"key2\":\"val2\"}"]
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

// paging extracts pagination data from the provided param and send it as a map.
//
//	This function expects 'paging' to be in the following format.
//	<url>?paging=["page:1","size:10"]
func paging(ctx *gin.Context) (map[string]uint32, error) {
	// default paging
	p := map[string]uint32{
		"page": 1,
		"size": 10,
	}

	paging := ctx.Query("paging")
	if paging == "" {
		return p, nil
	}

	var arr []string
	if err := json.Unmarshal([]byte(paging), &arr); err != nil {
		return p, errors.Join(errors.New("rest: error unmarshaling 'paging' from query param"), err)
	}

	for _, item := range arr {
		parts := strings.SplitN(item, ":", 2) // only split into 2 parts
		if len(parts) != 2 {
			return p, fmt.Errorf("rest: '%s' is of invalid format", item)
		}

		key := strings.TrimSpace(parts[0])
		val := strings.TrimSpace(parts[1])
		v, err := strconv.ParseUint(val, 10, 32)
		if err != nil {
			return p, fmt.Errorf("rest: paging '%s' is of invalid format", item)
		}

		p[key] = uint32(v)
	}

	return p, nil
}

// accountRequest maps account details sent in the request payload.
type accountRequest struct {
	Owner    string  `json:"owner" binding:"required"`
	Currency string  `json:"currency"`
	Balance  float32 `json:"balance"`
}
