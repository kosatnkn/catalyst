package transformers

import "github.com/kosatnkn/catalyst/domain/entities"

// SampleTransformer is used to transform sample
type SampleTransformer struct {
	ID   int64  `json:"id"`
	Name string `json:"name"`
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

	var tr []SampleTransformer

	for _, sample := range data.([]interface{}) {
		tr = append(tr, t.TransformAsObject(sample).(SampleTransformer))
	}

	return tr
}
