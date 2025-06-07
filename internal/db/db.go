package db

import (
	"database/sql"

	_ "github.com/mattn/go-sqlite3"
)

func NewDB() (*sql.DB, error) {
	db, err := sql.Open("sqlite3", "./data/myidentity.db")
	if err != nil {
		return nil, err
	}
	_, err = db.Exec(`
		CREATE TABLE IF NOT EXISTS users (
			id TEXT PRIMARY KEY,
			email TEXT NOT NULL UNIQUE,
			password TEXT NOT NULL,
			created_at DATETIME DEFAULT CURRENT_TIMESTAMP
		);
	`)
	return db, err
}
