package transformers

import (
	"github.com/kosatnkn/catalyst/v3/app/transport/http/response/transformer"
	"github.com/kosatnkn/catalyst/v3/internal/req/paginator"
)

// PaginatorTransformer is used to transform a paginator object.
type PaginatorTransformer struct {
	Page uint32 `json:"page"`
	Size uint8  `json:"size"`
}

// NewPaginatorTransformer creates a new instance of the transformer.
func NewPaginatorTransformer() transformer.TransformerInterface {
	return &PaginatorTransformer{}
}

// TransformAsObject map data to a transformer object.
func (t *PaginatorTransformer) TransformAsObject(data any) (any, error) {
	p, ok := data.(paginator.Paginator)
	if !ok {
		return nil, t.dataMismatchError()
	}

	tr := PaginatorTransformer{
		Page: p.Page,
		Size: uint8(p.Size),
	}

	return tr, nil
}

// TransformAsCollection map data to a collection of transformer objects.
func (t *PaginatorTransformer) TransformAsCollection(data any) (any, error) {
	return nil, transformer.NewTransformerError("", "Cannot transform paginator as a collection", nil)
}

// dataMismatchError returns a data mismatch error of TransformerError type.
func (t *PaginatorTransformer) dataMismatchError() error {
	return transformer.NewTransformerError("", "Cannot map given data to PaginatorTransformer", nil)
}
