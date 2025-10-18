package rest

import (
	"encoding/json"
	"errors"
	"fmt"
	"strings"
)

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
