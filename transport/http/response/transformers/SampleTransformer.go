package transformers

import (
	"github.com/kosatnkn/catalyst/v3/app/transport/http/response/transformer"
	"github.com/kosatnkn/catalyst/v3/domain/entities"
)

// SampleTransformer is used to transform sample
type SampleTransformer struct {
	ID   uint64 `json:"id"`
	Name string `json:"name"`
}

// NewSampleTransformer creates a new instance of the transformer.
func NewSampleTransformer() TransformerInterface {
	return &SampleTransformer{}
}

// TransformAsObject map data to a transformer object.
func (t *SampleTransformer) TransformAsObject(data interface{}) (interface{}, error) {
	sample, ok := data.(entities.Sample)
	if !ok {
		return nil, t.dataMismatchError()
	}

	tr := SampleTransformer{
		ID:   sample.ID,
		Name: sample.Name,
	}

	return tr, nil
}

// TransformAsCollection map data to a collection of transformer objects.
func (t *SampleTransformer) TransformAsCollection(data interface{}) (interface{}, error) {
	// Make sure that you declare the transformer slice in this manner.
	// Otherwise the marshaller will return `null` instead of `[]` when
	// marshalling empty slices
	// https://apoorvam.github.io/blog/2017/golang-json-marshal-slice-as-empty-array-not-null/
	trSamples := make([]SampleTransformer, 0)

	samples, ok := data.([]entities.Sample)
	if !ok {
		return nil, t.dataMismatchError()
	}

	for _, sample := range samples {
		tr, err := t.TransformAsObject(sample)
		if err != nil {
			return nil, err
		}

		trSample := tr.(SampleTransformer)
		trSamples = append(trSamples, trSample)
	}

	return trSamples, nil
}

// dataMismatchError returns a data mismatch error of TransformerError type.
func (t *SampleTransformer) dataMismatchError() error {
	return transformer.NewTransformerError("100", "Cannot map given data to SampleTransformer", nil)
}
