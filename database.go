package main

import (
	"database/sql"
	"log"
	"os"
	"path/filepath"

	_ "github.com/mattn/go-sqlite3"
)

func dbPath() string {
	homeDir, err := os.UserHomeDir()
	if err != nil {
		log.Fatalf("Error getting home directory: %v", err)
	}
	return filepath.Join(homeDir, "rauxy.db")
}

func OpenDB() (*sql.DB, error) {
	db, err := sql.Open("sqlite3", dbPath())
	if err != nil {
		return nil, err
	}

	createTableQuery := `
		CREATE TABLE IF NOT EXISTS tokens (
			id INTEGER PRIMARY KEY AUTOINCREMENT,
			name TEXT UNIQUE NOT NULL,
			token TEXT NOT NULL,
			port TEXT NOT NULL,
			created_at TEXT NOT NULL
		)
	`
	_, err = db.Exec(createTableQuery)
	if err != nil {
		log.Fatalf("Error creating table: %v\n", err)
	}
	return db, nil
}
