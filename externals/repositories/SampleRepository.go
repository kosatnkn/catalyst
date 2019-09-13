package repositories

import (
	"context"

	"github.com/kosatnkn/catalyst/domain/boundary/adapters"
	"github.com/kosatnkn/catalyst/domain/entities"
)

// SampleRepository is an example repository that implements test database functionality.
type SampleRepository struct {
	DBAdapter adapters.DBAdapterInterface
}

// Get temporarily mocks database functionality.
func (t *SampleRepository) Get(ctx context.Context) ([]entities.Sample, error) {

	// temporarily added
	var m []entities.Sample

	return m, nil

	// query := `SELECT id, name
	// 			FROM public.sample`

	// parameters := map[string]interface{}{
	// 	"id":   4,
	// 	"name": "Name 4",
	// }

	// result, err := t.DBAdapter.Query(query, parameters)
	// if err != nil {
	// 	return nil, err
	// }

	// return mapResult(result), nil
}

// Map results to entities.
func mapResult(result []map[string]interface{}) []entities.Sample {

	var m []entities.Sample

	for _, row := range result {

		id, _ := row["id"].(int64)
		name, _ := row["name"].(string)

		m = append(m, entities.Sample{
			ID:   id,
			Name: name,
		})
	}

	return m
}
