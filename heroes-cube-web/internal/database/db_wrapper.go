package database

import (
	"context"
	"database/sql"
	"fmt"

	_ "github.com/mattn/go-sqlite3"
)

// DBWrapper is a wrapper around sql.DB and Queries
type DB struct {
	*Queries
	db *sql.DB
}

// NewDBWrapper creates a new DB
func NewDBWrapper(db *sql.DB) *DB {
	return &DB{
		Queries: New(db),
		db:      db,
	}
}

// Open creates a new DB connection and returns a DB
func Open(dataSourceName string) (*DB, error) {
	db, err := sql.Open("sqlite3", dataSourceName)
	if err != nil {
		return nil, fmt.Errorf("error opening database: %w", err)
	}

	if err := db.Ping(); err != nil {
		return nil, fmt.Errorf("error connecting to database: %w", err)
	}

	return NewDBWrapper(db), nil
}

// Close closes the database connection
func (w *DB) Close() error {
	return w.db.Close()
}

// BeginTx starts a transaction
func (w *DB) BeginTx(ctx context.Context, opts *sql.TxOptions) (*sql.Tx, error) {
	return w.db.BeginTx(ctx, opts)
}

// WithTx creates a new Queries struct with a transaction
func (w *DB) WithTx(tx *sql.Tx) *Queries {
	return New(tx)
}
