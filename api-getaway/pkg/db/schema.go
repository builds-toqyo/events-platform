package db

import (
	"fmt"
	"github.com/jmoiron/sqlx"
)

var DB *sqlx.DB

func CreateTables() error {
	queries := []string{
		`CREATE TABLE IF NOT EXISTS users (
			id SERIAL PRIMARY KEY,
			name VARCHAR(255) NOT NULL,
			email VARCHAR(255) NOT NULL UNIQUE
		);`,

	}

	for _, query := range queries {
		_, err := DB.Exec(query)
		if err != nil {
			return fmt.Errorf("error executing query: %v", err)
		}
	}

	return nil
}
