package storage

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	_ "github.com/jackc/pgx/v5/stdlib"
)

var DB *sql.DB

// ConnectDB initializes the database connection
func ConnectDB() error {
	dbURL := os.Getenv("DATABASE_URL")
	if dbURL == "" {
		dbURL = "postgres://admin:adminpassword@postgres:5432/iam_db?sslmode=disable"
	}

	var err error
	DB, err = sql.Open("pgx", dbURL)
	if err != nil {
		return fmt.Errorf("failed to connect to DB: %w", err)
	}

	// Ping the database to verify connection
	if err := DB.Ping(); err != nil {
		return fmt.Errorf("DB ping failed: %w", err)
	}

	log.Println("Connected to PostgreSQL")
	return nil
}
