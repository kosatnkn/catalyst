package repositories

import (
	"context"

	"github.com/kosatnkn/catalyst/domain/boundary/adapters"
	"github.com/kosatnkn/catalyst/domain/boundary/repositories"
	"github.com/kosatnkn/catalyst/domain/entities"
)

// SampleMySQLRepository is an example repository that implements test database functionality.
type SampleMySQLRepository struct {
	db adapters.DBAdapterInterface
}

// NewSampleMySQLRepository creates a new instance of the repository.
func NewSampleMySQLRepository(dbAdapter adapters.DBAdapterInterface) repositories.SampleRepositoryInterface {

	return &SampleMySQLRepository{db: dbAdapter}
}

// Get retrieves a collection of Samples.
func (repo *SampleMySQLRepository) Get(ctx context.Context) ([]entities.Sample, error) {

	query := `SELECT id, name, password
				FROM sample`

	parameters := map[string]interface{}{}

	result, err := repo.db.Query(ctx, query, parameters)
	if err != nil {
		return nil, err
	}

	return repo.mapResult(result)
}

// GetByID retrieves a single Sample.
func (repo *SampleMySQLRepository) GetByID(ctx context.Context, id int) (entities.Sample, error) {

	// NOTE: DBAdapters in Catalyst supports named parameters and you don't have to
	// 		 worry about the order in which those parameters are declared in the
	//		 query and in the parameters map. The DBAdapter will take care of that.
	query := `SELECT id, name, password
				FROM sample
				WHERE id=?id`

	parameters := map[string]interface{}{
		"id": id,
	}

	result, err := repo.db.Query(ctx, query, parameters)
	if err != nil {
		return entities.Sample{}, err
	}

	mapped, err := repo.mapResult(result)
	if err != nil {
		return entities.Sample{}, err
	}

	if len(mapped) == 0 {
		return entities.Sample{}, nil
	}

	return mapped[0], nil
}

// Add adds a new sample.
func (repo *SampleMySQLRepository) Add(ctx context.Context, sample entities.Sample) error {

	query := `INSERT INTO sample
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
func (repo *SampleMySQLRepository) Edit(ctx context.Context, sample entities.Sample) error {

	query := `UPDATE sample
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
func (repo *SampleMySQLRepository) Delete(ctx context.Context, id int) error {

	query := `DELETE FROM sample
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
func (repo *SampleMySQLRepository) mapResult(result []map[string]interface{}) (samples []entities.Sample, err error) {

	// Applying type assertion in this manner will result in a panic when the db data structure changes.
	// This defer recover pattern is used to recover from the panic and to return an error instead.
	// Notice the use of `named returned values` for this function (without which the recover pattern will not work).
	defer func() {
		if r := recover(); r != nil {
			err = r.(error)
		}
	}()

	for _, row := range result {

		samples = append(samples, entities.Sample{
			ID:   int(row["id"].(int64)),
			Name: string(row["name"].([]byte)),
		})
	}

	return samples, nil
}
