package postgres

import (
	"context"
	"database/sql"
	"fmt"
	"regexp"
	"strings"

	// database driver for postgres
	_ "github.com/lib/pq"

	"github.com/kosatnkn/catalyst/v3/externals/adapters"
	"github.com/kosatnkn/catalyst/v3/internal/db"
)

// Adapter is used to communicate with a Postgres database.
type Adapter struct {
	cfg      Config
	pool     *sql.DB
	pqPrefix string
}

// NewAdapter creates a new Postgres adapter instance.
func NewAdapter(cfg Config) (adapters.DBAdapterInterface, error) {
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

	a := &Adapter{
		cfg:      cfg,
		pool:     db,
		pqPrefix: "?",
	}

	// check whether the db is accessible
	if cfg.Check {
		return a, a.Ping()
	}

	return a, nil
}

// Ping checks wether the database is accessible.
func (a *Adapter) Ping() error {
	return a.pool.Ping()
}

// Query runs a query and returns the result.
//
// Note: For INSERT statements postgres does not return the insert id by default.
// The returning identifier should be defined in the query using the RETURNING clause.
func (a *Adapter) Query(ctx context.Context, query string, params map[string]interface{}) ([]map[string]interface{}, error) {
	convertedQuery, placeholders := a.convertQuery(query)

	reorderedParams, err := a.reorderParameters(params, placeholders)
	if err != nil {
		return nil, err
	}

	stmt, err := a.prepareStatement(ctx, convertedQuery)
	if err != nil {
		return nil, err
	}
	defer stmt.Close()

	// check whether the query is a select statement
	if a.isSelect(convertedQuery) {
		rows, err := stmt.Query(reorderedParams...)
		if err != nil {
			return nil, err
		}

		return a.prepareDataSet(rows)
	}

	// check whether the query is an insert statement
	if a.isInsert(convertedQuery) {
		row := stmt.QueryRow(reorderedParams...)
		return a.prepareInsertResultSet(row)
	}

	result, err := stmt.Exec(reorderedParams...)
	// result, err := stmt.Query(reorderedParams...)
	if err != nil {
		return nil, err
	}

	return a.prepareResultSet(result)
}

// QueryBulk runs a query using an array of parameters and return the combined result.
//
// This query is intended to do bulk INSERTS, UPDATES and DELETES.
// Using this for SELECTS will result in an error.
func (a *Adapter) QueryBulk(ctx context.Context, query string, params []map[string]interface{}) ([]map[string]interface{}, error) {
	convertedQuery, placeholders := a.convertQuery(query)

	// check whether the query is a select statement
	if a.isSelect(convertedQuery) {
		return nil, fmt.Errorf("postgres-adapter: select queries are not allowed. use Query() instead")
	}

	stmt, err := a.prepareStatement(ctx, convertedQuery)
	if err != nil {
		return nil, err
	}
	defer stmt.Close()

	var lastID interface{}
	var affRows int64

	if a.isInsert(convertedQuery) {
		for _, pms := range params {
			reorderedParams, err := a.reorderParameters(pms, placeholders)
			if err != nil {
				return nil, err
			}

			row := stmt.QueryRow(reorderedParams...)
			if err := row.Err(); err != nil {
				return nil, err
			}

			row.Scan(&lastID)
			affRows++
		}

		return a.formatResultSet(lastID, affRows), nil
	}

	for _, pms := range params {
		reorderedParams, err := a.reorderParameters(pms, placeholders)
		if err != nil {
			return nil, err
		}

		result, err := stmt.Exec(reorderedParams...)
		if err != nil {
			return nil, err
		}

		ar, _ := result.RowsAffected()
		affRows += ar
	}

	return a.formatResultSet(lastID, affRows), nil
}

// WrapInTx runs the content of the function in a single transaction.
func (a *Adapter) WrapInTx(ctx context.Context, fn func(ctx context.Context) (interface{}, error)) (interface{}, error) {
	// attach a transaction to context
	ctx, err := a.attachTx(ctx)
	if err != nil {
		return nil, err
	}

	// get a reference to the attached transaction
	tx := ctx.Value(db.TxKey).(*sql.Tx)

	// run function
	res, err := fn(ctx)

	// decide whether to commit or rollback
	//
	// Here we deliberately avoid catching errors from Commit() and Rollback().
	// This is because the sql package does not give a method to check whether
	// a transaction has already completed or not.
	//
	// When executing nested operations in a single transaction, either the leaf operation or the
	// earliest failing operation of the operation tree will close the transaction.
	// Since all operations prior to that operation also tries to close the transaction
	// it will always result in an error.
	// If we catch errors from Commit() and Rollback(), nested transactions
	// will always fail because of this.
	if err != nil {
		tx.Rollback()
		return nil, err
	}

	tx.Commit()

	return res, nil
}

// Destruct will close the Postgres adapter releasing all resources.
func (a *Adapter) Destruct() error {
	return a.pool.Close()
}

// isSelect checks whether q is a select query.
func (a *Adapter) isSelect(q string) bool {
	return strings.ToLower(q[:6]) == "select"
}

// isInsert checks whether q is an insert query.
func (a *Adapter) isInsert(q string) bool {
	return strings.ToLower(q[:6]) == "insert"
}

// attachTx attaches a database transaction to the context.
//
// This will first check to see whether there is a transaction already in the context.
// Having a transaction already attached to context probably means that the calling function
// has been wrapped in a transaction in a previous stage.
// When this is the case use the existing attached transaction.
// Otherwise create a new transaction and attach.
func (a *Adapter) attachTx(ctx context.Context) (context.Context, error) {
	// check tx altready exists
	tx := ctx.Value(db.TxKey)
	if tx != nil {
		return ctx, nil
	}

	// attach new tx
	tx, err := a.pool.Begin()
	if err != nil {
		return nil, err
	}

	return context.WithValue(ctx, db.TxKey, tx), nil
}

// convertQuery converts the named parameter query to a placeholder query that Postgres library understands.
//
// Postgres placeholder formats look like this.
//
// SELECT * FROM tbl WHERE col = $1
// INSERT INTO tbl(col1, col2, col3) VALUES ($1, $2, $3)
// UPDATE tbl SET col1 = $1, col2 = $2 WHERE col3 = $3
// DELETE FROM tbl WHERE col = $1
//
// This will return the query and a slice of strings containing named parameter name in the order that they are found
// in the query.
func (a *Adapter) convertQuery(query string) (string, []string) {
	query = strings.TrimSpace(query)
	exp := regexp.MustCompile(`\` + a.pqPrefix + `\w+`)

	namedParams := exp.FindAllString(query, -1)

	for i := 0; i < len(namedParams); i++ {
		namedParams[i] = strings.TrimPrefix(namedParams[i], a.pqPrefix)
	}

	paramPosition := 0
	query = string(exp.ReplaceAllFunc([]byte(query), func(param []byte) []byte {
		paramPosition++
		paramName := fmt.Sprintf("$%d", paramPosition)

		return []byte(paramName)
	}))

	return query, namedParams
}

// reorderParameters reorders the parameters map in the order of named parameters slice.
func (a *Adapter) reorderParameters(params map[string]interface{}, namedParams []string) ([]interface{}, error) {
	var reorderedParams []interface{}

	for _, param := range namedParams {
		// return an error if a named parameter is missing from params
		paramValue, ok := params[param]
		if !ok {
			return nil, fmt.Errorf("postgres-adapter: parameter '%s' is missing", param)
		}

		reorderedParams = append(reorderedParams, paramValue)
	}

	return reorderedParams, nil
}

// prepareStatement creates a prepared statement using the query.
//
// Checks whether there is a transaction attached to the context.
// If so use that transaction to prepare statement else use the pool.
func (a *Adapter) prepareStatement(ctx context.Context, query string) (*sql.Stmt, error) {
	tx := ctx.Value(db.TxKey)
	if tx != nil {
		return tx.(*sql.Tx).Prepare(query)
	}

	return a.pool.Prepare(query)
}

// prepareDataSet creates a dataset using the output of a SELECT statement.
//
// Source: https://kylewbanks.com/blog/query-result-to-map-in-golang
func (a *Adapter) prepareDataSet(rows *sql.Rows) ([]map[string]interface{}, error) {
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

// prepareInsertResultSet creates a resultset using the result of QueryRow().
func (a *Adapter) prepareInsertResultSet(row *sql.Row) ([]map[string]interface{}, error) {
	if err := row.Err(); err != nil {
		return nil, err
	}

	var id interface{}
	row.Scan(&id)

	return a.formatResultSet(id, 1), nil
}

// prepareResultSet creates a resultset using the result of Exec().
//
// This is used with UPDATE and DELETE statements.
//
// Note: result.LastInsertId() is not supported by Postgres. So this cannot be used with INSERT statements.
func (a *Adapter) prepareResultSet(result sql.Result) ([]map[string]interface{}, error) {
	aff, err := result.RowsAffected()
	if err != nil {
		return nil, err
	}

	return a.formatResultSet(nil, aff), nil
}

// formatResultSet creates a resultset using last insert id and affected rows.
func (a *Adapter) formatResultSet(id interface{}, aff int64) []map[string]interface{} {
	data := make([]map[string]interface{}, 0)

	return append(data, map[string]interface{}{
		db.AffectedRows: aff,
		db.LastInsertID: id,
	})
}
