package adapters

import (
	"context"
	"database/sql"
	"fmt"
	"regexp"
	"strings"

	// database driver for postgres
	_ "github.com/lib/pq"

	"github.com/kosatnkn/catalyst/app/config"
	errTypes "github.com/kosatnkn/catalyst/app/error/types"
	"github.com/kosatnkn/catalyst/domain/boundary/adapters"
	"github.com/kosatnkn/catalyst/domain/globals"
)

// MariaDBAdapter is used to communicate with a Postgres database.
type MariaDBAdapter struct {
	cfg      config.DBConfig
	pool     *sql.DB
	pqPrefix string
}

// NewPostgresAdapter creates a new Postgres adapter instance.
func NewMariaDBAdapter(cfg config.DBConfig) (adapters.DBAdapterInterface, error) {

	connString := fmt.Sprintf("user=%s password=%s dbname=%s host=%s port=%d sslmode=disable",
		cfg.User, cfg.Password, cfg.Database, cfg.Host, cfg.Port)

	db, err := sql.Open("postgres", connString)
	if err != nil {
		return nil, err
	}

	// pool configurations
	db.SetMaxOpenConns(cfg.PoolSize)
	//db.SetMaxIdleConns(2)
	//db.SetConnMaxLifetime(time.Hour)

	a := &MariaDBAdapter{
		cfg:      cfg,
		pool:     db,
		pqPrefix: "?",
	}

	return a, nil
}

// Query runs a query and returns the result.
func (a *MariaDBAdapter) Query(ctx context.Context, query string, parameters map[string]interface{}) ([]map[string]interface{}, error) {

	convertedQuery, placeholders := a.convertQuery(query)

	reorderedParameters, err := a.reorderParameters(parameters, placeholders)
	if err != nil {
		return nil, err
	}

	statement, err := a.prepareStatement(ctx, convertedQuery)
	if err != nil {
		return nil, err
	}

	// check whether the query is a select statement
	if strings.ToLower(convertedQuery[:1]) == "s" {

		rows, err := statement.Query(reorderedParameters...)
		if err != nil {
			return nil, err
		}

		return a.prepareDataSet(rows)
	}

	result, err := statement.Exec(reorderedParameters...)
	if err != nil {
		return nil, err
	}

	return a.prepareResultSet(result)
}

// NewTransaction creates a new database transaction.
func (a *MariaDBAdapter) NewTransaction() (*sql.Tx, error) {

	return a.pool.Begin()
}

// Destruct will close the Postgres adapter releasing all resources.
func (a *MariaDBAdapter) Destruct() {

	a.pool.Close()
}

// Convert the named parameter query to a placeholder query that Postgres library understands.
// This will return the query and a slice of strings containing named parameter name in the order that they are found
// in the query.
func (a *MariaDBAdapter) convertQuery(query string) (string, []string) {

	query = strings.TrimSpace(query)
	exp := regexp.MustCompile(`\` + a.pqPrefix + `\w+`)

	namedParams := exp.FindAllString(query, -1)

	for i := 0; i < len(namedParams); i++ {
		namedParams[i] = strings.TrimPrefix(namedParams[i], a.pqPrefix)
	}

	paramPosition := 0
	//query = string(exp.ReplaceAllFunc([]byte(query), a.replaceNamedParam))
	query = string(exp.ReplaceAllFunc([]byte(query), func(param []byte) []byte {

		paramPosition++
		paramName := fmt.Sprintf("$%d", paramPosition)

		return []byte(paramName)
	}))

	return query, namedParams
}

// Reorder the parameters map in the order of named parameters slice.
func (a *MariaDBAdapter) reorderParameters(params map[string]interface{}, namedParams []string) ([]interface{}, error) {

	var reorderedParams []interface{}

	for _, param := range namedParams {

		// return an error if a named parameter is missing from params
		paramValue, isParamExist := params[param]

		if !isParamExist {
			return nil, errTypes.NewAdapterError(fmt.Sprintf("parameter '%s' is missing", param), 100, "")
		}

		reorderedParams = append(reorderedParams, paramValue)
	}

	return reorderedParams, nil
}

// prepareStatement creates a prepared statement using the query.
//
// Checks whether there is a transaction attached to the context.
// If so use that transaction to prepare statement else use the pool.
func (a *MariaDBAdapter) prepareStatement(ctx context.Context, query string) (*sql.Stmt, error) {

	tx := ctx.Value(globals.TxKey)
	if tx != nil {
		return tx.(*sql.Tx).Prepare(query)
	}

	return a.pool.Prepare(query)
}

// Prepare the return dataset for select statements.
// Source: https://kylewbanks.com/blog/query-result-to-map-in-golang
func (a *MariaDBAdapter) prepareDataSet(rows *sql.Rows) ([]map[string]interface{}, error) {

	defer rows.Close()

	var data []map[string]interface{}
	cols, _ := rows.Columns()

	// create a slice of interface{}'s to represent each column
	// and a second slice to contain pointers to each item in the columns slice
	columns := make([]interface{}, len(cols))
	columnPointers := make([]interface{}, len(cols))

	for i := range columns {
		columnPointers[i] = &columns[i]
	}

	for rows.Next() {
		// scan the result into the column pointers
		err := rows.Scan(columnPointers...)
		if err != nil {
			return nil, err
		}

		// create our map, and retrieve the value for each column from the pointers slice
		// storing it in the map with the name of the column as the key
		row := make(map[string]interface{})

		for i, colName := range cols {
			val := columnPointers[i].(*interface{})
			row[colName] = *val
		}

		data = append(data, row)
	}

	return data, nil
}

// Prepare the resultset for all other queries.
func (a *MariaDBAdapter) prepareResultSet(result sql.Result) ([]map[string]interface{}, error) {

	var data []map[string]interface{}

	row := make(map[string]interface{})

	row["affected_rows"], _ = result.RowsAffected()
	row["last_insert_id"], _ = result.LastInsertId()

	return append(data, row), nil
}
