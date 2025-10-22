package postgres

import "fmt"

func withPagination(query string, paging map[string]uint32) string {
	if paging["page"] == 0 {
		paging["page"] = 1
	}

	return fmt.Sprintf("%s LIMIT %d OFFSET %d", query, paging["size"], (paging["page"]-1)*paging["size"])
}

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
