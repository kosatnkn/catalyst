package repositories

import (
	"github.com/kosatnkn/catalyst/domain/boundary/adapters"
	"github.com/kosatnkn/catalyst/domain/entities"
)

// TestRepository is an example repository that implements test database functionality.
type TestRepository struct {
	DBAdapter adapters.DBAdapterInterface
}

// TestRepositoryMethod tests database functionality.
func (t *TestRepository) TestRepositoryMethod() ([]entities.Test, error) {

	query := `SELECT id, name
				FROM public.test`

	parameters := map[string]interface{}{
		"id":   4,
		"name": "Name 4",
	}

	result, err := t.DBAdapter.Query(query, parameters)

	if err != nil {
		return nil, err
	}

	return mapResult(result), nil
}

// Map results to entities.
func mapResult(result []map[string]interface{}) []entities.Test {

	var m []entities.Test

	for _, row := range result {

		id, _ := row["id"].(int64)
		name, _ := row["name"].(string)

		m = append(m, entities.Test{
			ID:   id,
			Name: name,
		})
	}

	return m
}
