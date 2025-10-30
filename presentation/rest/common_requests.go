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
func filters(param string) (map[string]any, error) {
	f := map[string]any{}
	if param == "" {
		return f, nil
	}

	var arr []string
	if err := json.Unmarshal([]byte(param), &arr); err != nil {
		return f, errors.Join(errors.New("rest: error extracting filters from query param"), err)
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
