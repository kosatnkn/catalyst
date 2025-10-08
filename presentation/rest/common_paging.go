package rest

import (
	"encoding/json"
	"errors"
	"fmt"
	"strconv"
	"strings"
)

func paging(param string) (map[string]uint32, error) {
	p := map[string]uint32{
		"page": 1,
		"size": 10,
	}
	if param == "" {
		return p, nil
	}

	var arr []string
	if err := json.Unmarshal([]byte(param), &arr); err != nil {
		return p, errors.Join(errors.New("rest: error extracting 'paging' from query param"), err)
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
