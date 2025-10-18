package postgres

import "fmt"

func withPagination(query string, paging map[string]uint32) string {
	if paging["page"] == 0 {
		paging["page"] = 1
	}

	return fmt.Sprintf("%s LIMIT %d OFFSET %d", query, paging["size"], (paging["page"]-1)*paging["size"])
}
