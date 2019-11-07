package repositories

import (
	"context"
	"fmt"

	"github.com/kosatnkn/catalyst/domain/boundary/adapters"
	"github.com/kosatnkn/catalyst/domain/boundary/repositories"
	"github.com/kosatnkn/catalyst/domain/entities"
)

// SampleRepository is an example repository that implements test database functionality.
type SampleRepository struct {
	db adapters.DBAdapterInterface
}

// NewSampleRepository creates a new instance of the repository.
func NewSampleRepository(dbAdapter adapters.DBAdapterInterface) repositories.SampleRepositoryInterface {

	return &SampleRepository{db: dbAdapter}
}

// Get retrieves a collection of Samples.
func (repo *SampleRepository) Get(ctx context.Context) ([]entities.Sample, error) {

	query := `SELECT id, name
				FROM public.sample`

	parameters := map[string]interface{}{
		"id":   4,
		"name": "Name 4",
	}

	result, err := repo.db.Query(ctx, query, parameters)
	if err != nil {
		return nil, err
	}

	return mapResult(result), nil
}

// GetByID retrieves a single Sample.
func (repo *SampleRepository) GetByID(ctx context.Context, id int) (entities.Sample, error) {

	return entities.Sample{
		ID:       int64(id),
		Name:     fmt.Sprintf("Name %d", id),
		Password: fmt.Sprintf("Password %d", id),
	}, nil
}

// Add adds a new sample.
func (repo *SampleRepository) Add(ctx context.Context, sample entities.Sample) error {

	return nil
}

// Edit updates an existing sample identified by the id.
func (repo *SampleRepository) Edit(ctx context.Context, id int, sample entities.Sample) error {

	return nil
}

// Delete deletes an existing sample identified by id.
func (repo *SampleRepository) Delete(ctx context.Context, id int) error {

	return nil
}

// mapResult maps the result to entities.
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
