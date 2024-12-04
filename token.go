package main

import (
	"crypto/rand"
	"encoding/hex"
	"fmt"
	"time"
)

func generateToken() (string, error) {
	bytes := make([]byte, 16)
	if _, err := rand.Read(bytes); err != nil {
		return "", err
	}
	return hex.EncodeToString(bytes), nil
}

func AddToken(name, port string) error {
	db, err := OpenDB()
	if err != nil {
		return err
	}
	defer db.Close()

	token, err := generateToken()
	if err != nil {
		return err
	}

	query := `
		INSERT INTO tokens (name, token, port, created_at)
		VALUES (?, ?, ?, ?)
	`
	_, err = db.Exec(query, name, token, port, time.Now())
	return err
}

func RemoveToken(name string) error {
	db, err := OpenDB()
	if err != nil {
		return err
	}
	defer db.Close()

	query := `DELETE FROM tokens WHERE name = ?`
	_, err = db.Exec(query, name)
	return err
}

func ListTokens() error {
	db, err := OpenDB()
	if err != nil {
		return err
	}
	defer db.Close()

	query := `SELECT id, name, token, port, created_at FROM tokens`
	rows, err := db.Query(query)
	if err != nil {
		return err
	}
	defer rows.Close()

	fmt.Printf("%-10s %-20s %-40s %-10s %-20s\n", "ID", "Name", "Token", "Port", "Created At")
	for rows.Next() {
		var id int
		var name, token, port, createdAt string
		if err := rows.Scan(&id, &name, &token, &port, &createdAt); err != nil {
			return err
		}
		fmt.Printf("%-10d %-20s %-40s %-10s %-20s\n", id, name, token, port, createdAt)
	}
	return nil
}
