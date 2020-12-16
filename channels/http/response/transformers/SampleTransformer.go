package transformers

import (
	"errors"

	"github.com/kosatnkn/catalyst/domain/entities"
)

// SampleTransformer is used to transform sample
type SampleTransformer struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

// NewSampleTransformer creates a new instance of the transformer.
func NewSampleTransformer() TransformerInterface {
	return &SampleTransformer{}
}

// TransformAsObject map data to a transformer object.
func (t *SampleTransformer) TransformAsObject(data interface{}) (interface{}, error) {

	// NOTE: Make sure to do the type assertion in this manner so that assertion failures
	// 		 will not result in a panic.
	// https://tour.golang.org/methods/15
	sample, ok := data.(entities.Sample)
	if !ok {
		return nil, errors.New("Unmatched data type")
	}

	tr := SampleTransformer{
		ID:   sample.ID,
		Name: sample.Name,
	}

	return tr, nil
}

// TransformAsCollection map data to a collection of transformer objects.
func (t *SampleTransformer) TransformAsCollection(data interface{}) (interface{}, error) {

	// NOTE: Make sure that you declare the transformer slice in this manner.
	//		 Otherwise the marshaller will return `null` instead of `[]` when
	//		 marshalling empty slices
	// https://apoorvam.github.io/blog/2017/golang-json-marshal-slice-as-empty-array-not-null/
	trSamples := make([]SampleTransformer, 0)

	samples, ok := data.([]entities.Sample)
	if !ok {
		return nil, errors.New("Unmatched data type")
	}

	for _, sample := range samples {

		tr, err := t.TransformAsObject(sample)
		if err != nil {
			return nil, err
		}

		trSample, ok := tr.(SampleTransformer)
		if !ok {
			return nil, errors.New("Unmatched data type")
		}

		trSamples = append(trSamples, trSample)
	}

	return trSamples, nil
}
