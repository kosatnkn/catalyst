package transformers

import (
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
func (t *SampleTransformer) TransformAsObject(data interface{}) interface{} {

	var sample = data.(entities.Sample)

	return SampleTransformer{
		ID:   sample.ID,
		Name: sample.Name,
	}

}

// TransformAsCollection map data to a collection of transformer objects.
func (t *SampleTransformer) TransformAsCollection(data interface{}) interface{} {

	// NOTE: Make sure that you declare the transformer slice in this manner.
	//		 Otherwise the marshaller will return `null` instead of `[]` when
	//		 marshalling empty slices
	// https://apoorvam.github.io/blog/2017/golang-json-marshal-slice-as-empty-array-not-null/
	tr := make([]SampleTransformer, 0)

	for _, sample := range data.([]entities.Sample) {
		tr = append(tr, t.TransformAsObject(sample).(SampleTransformer))
	}

	return tr
}
