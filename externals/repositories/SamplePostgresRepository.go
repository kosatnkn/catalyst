package repositories

import (
	"context"

	"github.com/kosatnkn/catalyst/domain/boundary/adapters"
	"github.com/kosatnkn/catalyst/domain/boundary/repositories"
	"github.com/kosatnkn/catalyst/domain/entities"
)

// SamplePostgresRepository is an example repository that implements test database functionality.
type SamplePostgresRepository struct {
	db adapters.DBAdapterInterface
}

// NewSampleMySQLRepository creates a new instance of the repository.
func NewSamplePostgresRepository(dbAdapter adapters.DBAdapterInterface) repositories.SampleRepositoryInterface {

	return &SamplePostgresRepository{db: dbAdapter}
}

// Get retrieves a collection of Samples.
func (repo *SamplePostgresRepository) Get(ctx context.Context) ([]entities.Sample, error) {

	query := `SELECT "id", "name", "password"
				FROM test.sample`

	parameters := map[string]interface{}{}

	result, err := repo.db.Query(ctx, query, parameters)
	if err != nil {
		return nil, err
	}

	return repo.mapResult(result), nil
}

// GetByID retrieves a single Sample.
func (repo *SamplePostgresRepository) GetByID(ctx context.Context, id int) (entities.Sample, error) {

	// NOTE: DBAdapters in Catalyst supports named parameters and you don't have to
	// 		 worry about the order in which those parameters are declared in the
	//		 query and in the parameters map. The DBAdapter will take care of that.
	query := `SELECT "id", "name", "password"
				FROM test.sample
				WHERE "id"=?id`

	parameters := map[string]interface{}{
		"id": id,
	}

	result, err := repo.db.Query(ctx, query, parameters)
	if err != nil {
		return entities.Sample{}, err
	}

	mapped := repo.mapResult(result)
	if len(mapped) == 0 {
		return entities.Sample{}, nil
	}

	return mapped[0], nil
}

// Add adds a new sample.
func (repo *SamplePostgresRepository) Add(ctx context.Context, sample entities.Sample) error {

	query := `INSERT INTO test.sample
				("name", "password")
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
func (repo *SamplePostgresRepository) Edit(ctx context.Context, sample entities.Sample) error {

	query := `UPDATE test.sample
				SET "name"=?name, "password"=?password
				WHERE "id"=?id
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
func (repo *SamplePostgresRepository) Delete(ctx context.Context, id int) error {

	query := `DELETE FROM test.sample
				WHERE "id"=?id
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
func (repo *SamplePostgresRepository) mapResult(result []map[string]interface{}) []entities.Sample {

	var m []entities.Sample

	for _, row := range result {

		id, _ := row["id"].(int64)
		name, _ := row["name"].([]byte)

		m = append(m, entities.Sample{
			ID:   int(id),
			Name: string(name),
		})
	}

	return m
}
