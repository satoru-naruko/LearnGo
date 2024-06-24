package my_database

import (
	"database/sql"
	"fmt"

	_ "github.com/mattn/go-sqlite3"
)

type DB struct {
	*sql.DB
}

func New(dsn string) (*DB, error) {
	db, err := sql.Open("sqlite3", dsn)
	if err != nil {
		return nil, fmt.Errorf("failed to open database: %w", err)
	}
	return &DB{db}, nil
}

func (db *DB) Query(query string, args ...interface{}) (*sql.Rows, error) {
	rows, err := db.DB.Query(query, args...)
	if err != nil {
		return nil, fmt.Errorf("failed to execute query: %w", err)
	}
	return rows, nil
}

func (db *DB) Exec(query string, args ...interface{}) (sql.Result, error) {
	res, err := db.DB.Exec(query, args...)
	if err != nil {
		return nil, fmt.Errorf("failed to execute query: %w", err)
	}
	return res, nil
}
