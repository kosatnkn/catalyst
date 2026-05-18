package postgres

import (
	"fmt"
	"regexp"
)

// Identity provides a single reference point to be used
// as an identifier for package resources.
const Identity = "postgres"

// withPagination adds pagination to the query.
func withPagination(query string, paging map[string]uint32) string {
	if paging["page"] == 0 {
		paging["page"] = 1
	}

	return fmt.Sprintf("%s LIMIT %d OFFSET %d", query, paging["size"], (paging["page"]-1)*paging["size"])
}

// allowedFiltersOnly returns the subset of filters that is allowed to be used with queries.
// This is to prevent injection of unknown filters.
//
// Deprecated: Using direct check to determine whether the key is allowed.
func allowedFiltersOnly(filters map[string]any, allowedKeys []string) map[string]any {
	// initialize the result map
	fts := make(map[string]any)

	// build a lookup set for O(1) key checks
	// NOTE: Converting the slice in to a map reduces the key check complexity from O(n) per check to O(1) per check.
	// `struct{}` takes zero bytes of memory, so `map[string]struct{}` is memory-efficient.
	allowed := make(map[string]struct{}, len(allowedKeys))
	for _, k := range allowedKeys {
		allowed[k] = struct{}{}
	}

	// copy only allowed keys
	for k, v := range filters {
		if _, ok := allowed[k]; ok {
			fts[k] = v
		}
	}

	return fts
}

// appendToQuery appends the queryPart to the query using the joiner.
func appendToQuery(query, joiner, queryPart string) (string, error) {
	// the joiner should be either `AND` or `OR`
	if !(joiner == `AND` || joiner == `OR`) {
		return "", fmt.Errorf(`postgres-helper-appendtoquery: joiner should be either 'AND' or 'OR'`)
	}

	// if query has a where append joiner + ` ` + queryPart
	regex := regexp.MustCompile(`(?i)\bWHERE\b`)
	if regex.MatchString(query) {
		return fmt.Sprintf(`%s %s %s`, query, joiner, queryPart), nil
	}

	return fmt.Sprintf(`%s WHERE %s`, query, queryPart), nil
}

// countQuery returns the query wrapped in a count query.
func countQuery(query string) string {
	return fmt.Sprintf(`SELECT COUNT(*) AS c FROM (%s) AS q`, query)
}

// to return in converted to T
func to[T any](in any, idx int) (T, error) {
	out, ok := in.(T)
	if !ok {
		return out, fmt.Errorf("postgres-helper-to: row (%d) invalid data type %T", idx, in)
	}

	return out, nil
}
