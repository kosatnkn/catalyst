package repositories

import (
	"context"

	"github.com/kosatnkn/catalyst/domain/boundary/adapters"
	"github.com/kosatnkn/catalyst/domain/boundary/repositories"
	"github.com/kosatnkn/catalyst/domain/entities"
)

// SampleRepository is an example repository that implements test database functionality.
type SampleRepository struct {
	db adapters.DBAdapterInterface
}

// NewSampleRepository creates a new instance of the repository
func NewSampleRepository(dbAdapter adapters.DBAdapterInterface) repositories.SampleRepositoryInterface {

	return &SampleRepository{db: dbAdapter}
}

// Get temporarily mocks database functionality.
func (repo *SampleRepository) Get(ctx context.Context) ([]entities.Sample, error) {

	// temporarily added
	var m []entities.Sample

	return m, nil

	// query := `SELECT id, name
	// 			FROM public.sample`

	// parameters := map[string]interface{}{
	// 	"id":   4,
	// 	"name": "Name 4",
	// }

	// result, err := repo.db.Query(query, parameters)
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
