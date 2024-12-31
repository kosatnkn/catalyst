package paginator

import (
	"fmt"
)

// PaginatorRepositoryFacilitator is the facilitator that will add pagination handling capabilities to the repository.
type PaginatorRepositoryFacilitator struct{}

// NewPaginatorRepositoryFacilitator creates a new instance of the facilitator.
func NewPaginatorRepositoryFacilitator() *PaginatorRepositoryFacilitator {
	return &PaginatorRepositoryFacilitator{}
}

// WithPagination attaches the pagination clause to the query.
func (repo *PaginatorRepositoryFacilitator) WithPagination(query string, p Paginator) string {
	return fmt.Sprintf("%s limit %d offset %d", query, p.Size, (p.Page-1)*p.Size)
}
