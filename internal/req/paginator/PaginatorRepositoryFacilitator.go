package paginator

import (
	"fmt"

	"github.com/kosatnkn/catalyst/v3/internal/req"
)

// PaginatorRepositoryFacilitator is the facilitator that will add pagination handling capabilities to the repository.
type PaginatorRepositoryFacilitator struct {
	dbType string
}

// NewPaginatorRepositoryFacilitator creates a new instance of the facilitator.
func NewPaginatorRepositoryFacilitator(dbType string) *PaginatorRepositoryFacilitator {
	return &PaginatorRepositoryFacilitator{
		dbType: dbType,
	}
}

// WithPagination attaches the pagination clause to the query.
func (repo *PaginatorRepositoryFacilitator) WithPagination(query string, p Paginator) string {
	switch repo.dbType {
	case req.DBMySQL:
		return fmt.Sprintf("%s LIMIT %d OFFSET %d", query, p.Size, (p.Page-1)*p.Size)
	case req.DBPostgres:
		return fmt.Sprintf("%s limit %d offset %d", query, p.Size, (p.Page-1)*p.Size)
	default:
		return ""
	}
}
