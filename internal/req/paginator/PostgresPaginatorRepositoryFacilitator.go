package paginator

import (
	"fmt"
)

// PostgresPaginatorRepositoryFacilitator is the facilitator that will add pagination handling capabilities to the repository.
type PostgresPaginatorRepositoryFacilitator struct{}

// NewPostgresPaginatorRepositoryFacilitator creates a new instance of the facilitator.
func NewPostgresPaginatorRepositoryFacilitator() *PostgresPaginatorRepositoryFacilitator {
	return &PostgresPaginatorRepositoryFacilitator{}
}

// WithPagination attaches the pagination clause to the query.
func (repo *PostgresPaginatorRepositoryFacilitator) WithPagination(query string, p Paginator) string {
	return fmt.Sprintf("%s limit %d offset %d", query, p.Size, (p.Page-1)*p.Size)
}
