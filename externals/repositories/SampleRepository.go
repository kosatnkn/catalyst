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

// NewSampleRepository creates a new instance of the repository.
func NewSampleRepository(dbAdapter adapters.DBAdapterInterface) repositories.SampleRepositoryInterface {

	return &SampleRepository{db: dbAdapter}
}

// Get retrieves a collection of Samples.
func (repo *SampleRepository) Get(ctx context.Context) ([]entities.Sample, error) {

	query := `SELECT id, name, password
				FROM test.sample`

	parameters := map[string]interface{}{}

	result, err := repo.db.Query(ctx, query, parameters)
	if err != nil {
		return nil, err
	}

	return mapResult(result), nil
}

// GetByID retrieves a single Sample.
func (repo *SampleRepository) GetByID(ctx context.Context, id int) (entities.Sample, error) {

	query := `SELECT id, name, password
				FROM test.sample
				WHERE id=?id`

	parameters := map[string]interface{}{
		"id": id,
	}

	result, err := repo.db.Query(ctx, query, parameters)
	if err != nil {
		return entities.Sample{}, err
	}

	mapped := mapResult(result)
	if len(mapped) == 0 {
		return entities.Sample{}, nil
	}

	return mapped[0], nil
}

// Add adds a new sample.
func (repo *SampleRepository) Add(ctx context.Context, sample entities.Sample) error {

	query := `INSERT INTO test.sample
				(name, password)
				VALUES(?name, ?password)
				`

	parameters := map[string]interface{}{
		"name":     sample.Name,
		"password": sample.Password,
	}

	_, err := repo.db.Query(ctx, query, parameters)
	if err != nil {
		return err
	}

	return nil
}

// Edit updates an existing sample identified by the id.
func (repo *SampleRepository) Edit(ctx context.Context, sample entities.Sample) error {

	query := `UPDATE test.sample
				SET name=?name, password=?password
				WHERE id=?id
				`

	parameters := map[string]interface{}{
		"id":       sample.ID,
		"name":     sample.Name,
		"password": sample.Password,
	}

	_, err := repo.db.Query(ctx, query, parameters)
	if err != nil {
		return err
	}

	return nil
}

// Delete deletes an existing sample identified by id.
func (repo *SampleRepository) Delete(ctx context.Context, id int) error {

	query := `DELETE FROM test.sample
				WHERE id=?id
				`

	parameters := map[string]interface{}{
		"id": id,
	}

	_, err := repo.db.Query(ctx, query, parameters)
	if err != nil {
		return err
	}

	return nil
}

// mapResult maps the result to entities.
func mapResult(result []map[string]interface{}) []entities.Sample {

	var m []entities.Sample

	for _, row := range result {

		id, _ := row["id"].(int64)
		name, _ := row["name"].(string)

		m = append(m, entities.Sample{
			ID:   int(id),
			Name: name,
		})
	}

	return m
}
