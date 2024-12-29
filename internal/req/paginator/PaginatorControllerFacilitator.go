package paginator

// PaginatorControllerFacilitator is the facilitator that will add pagination handling capabilities to the controller.
type PaginatorControllerFacilitator struct{}

// NewPaginatorControllerFacilitator creates a new instance of the facilitator.
func NewPaginatorControllerFacilitator() *PaginatorControllerFacilitator {
	return &PaginatorControllerFacilitator{}
}

// GetPaginator extracts pagination data from query parameters
func (ctl *PaginatorControllerFacilitator) GetPaginator(page, size uint32) Paginator {
	return Paginator{
		Page: page,
		Size: size,
	}
}
