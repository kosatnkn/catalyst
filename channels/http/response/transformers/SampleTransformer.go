package transformers

import (
	"github.com/kosatnkn/catalyst/channels/http/errors"
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
func (t *SampleTransformer) TransformAsObject(data interface{}) (_ interface{}, err error) {

	// NOTE: Applying type assertion in this manner will result in a panic if the assertion fails.
	// 		 This defer recover pattern is used to recover from the panic and to return an error instead.
	// 		 Notice the use of `named returned values` for this function (without which the recover pattern will not work).
	defer func() {
		if r := recover(); r != nil {
			err = errors.NewTransformerError(r.(error).Error(), 100, "")
		}
	}()

	sample := data.(entities.Sample)

	tr := SampleTransformer{
		ID:   sample.ID,
		Name: sample.Name,
	}

	return tr, nil
}

// TransformAsCollection map data to a collection of transformer objects.
func (t *SampleTransformer) TransformAsCollection(data interface{}) (_ interface{}, err error) {

	// NOTE: Make sure that you declare the transformer slice in this manner.
	//		 Otherwise the marshaller will return `null` instead of `[]` when
	//		 marshalling empty slices
	// https://apoorvam.github.io/blog/2017/golang-json-marshal-slice-as-empty-array-not-null/
	trSamples := make([]SampleTransformer, 0)

	// NOTE: Applying type assertion in this manner will result in a panic if the assertion fails.
	// 		 This defer recover pattern is used to recover from the panic and to return an error instead.
	// 		 Notice the use of `named returned values` for this function (without which the recover pattern will not work).
	defer func() {
		if r := recover(); r != nil {
			err = errors.NewTransformerError(r.(error).Error(), 100, "")
		}
	}()

	samples := data.([]entities.Sample)

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
