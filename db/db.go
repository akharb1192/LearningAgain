// db/db.go
package db

import "database/sql"

// DBInterface defines the methods for interacting with a database.
type DBInterface interface {
	// Exec executes a query that doesn't return rows (e.g., INSERT, UPDATE, DELETE)
	Exec(query string, args ...interface{}) (ResultInterface, error)
}

// SQLDBWrapper is a wrapper around *sql.DB to implement DBInterface
type SQLDBWrapper struct {
	DB *sql.DB
}

// Exec executes a query that doesn't return rows (e.g., INSERT, UPDATE, DELETE)
func (wrapper *SQLDBWrapper) Exec(query string, args ...interface{}) (ResultInterface, error) {
	result, err := wrapper.DB.Exec(query, args...)
	if err != nil {
		return nil, err
	}
	return &SQLResultWrapper{Result: result}, nil
}

// SQLResultWrapper is a wrapper around sql.Result to implement ResultInterface
type SQLResultWrapper struct {
	Result sql.Result
}

// LastInsertId returns the last inserted row ID
func (wrapper *SQLResultWrapper) LastInsertId() (int64, error) {
	return wrapper.Result.LastInsertId()
}

// ResultInterface defines the methods for results from non-SELECT queries (like INSERT/UPDATE/DELETE)
type ResultInterface interface {
	LastInsertId() (int64, error)
}
