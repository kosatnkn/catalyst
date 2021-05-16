package unpackers

// PaginatorUnpacker contains the unpacking structure for paginator request payload.
//
// https://pkg.go.dev/gopkg.in/go-playground/validator.v10#section-documentation
type PaginatorUnpacker struct {
	Page uint32 `json:"page" validate:"min=1"`
	Size uint32 `json:"size" validate:"min=1,max=100"`
}

// NewPaginatorUnpacker creates a new instance of the unpacker.
func NewPaginatorUnpacker() *PaginatorUnpacker {
	return &PaginatorUnpacker{}
}

// RequiredFormat returns the applicable JSON format for the address data structure.
func (u *PaginatorUnpacker) RequiredFormat() string {
	return `{
		"page": "<integer, required, min=1>",
		"size": "<integer, required, min=1, max=100>"
	}`
}
