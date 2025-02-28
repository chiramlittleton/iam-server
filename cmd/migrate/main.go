package main

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	_ "github.com/jackc/pgx/v5/stdlib"
)

func main() {
	dbURL := os.Getenv("DATABASE_URL")
	if dbURL == "" {
		dbURL = "postgres://admin:adminpassword@postgres:5432/iam_db?sslmode=disable"
	}

	db, err := sql.Open("pgx", dbURL) // ✅ Correct driver name for pgx
	if err != nil {
		log.Fatal("Cannot connect to DB:", err)
	}
	defer db.Close()

	migrations := []string{
		`CREATE TABLE IF NOT EXISTS users (
			id SERIAL PRIMARY KEY,
			email VARCHAR(255) UNIQUE NOT NULL,
			password_hash TEXT NOT NULL,
			created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
		);`,
	}

	for _, migration := range migrations {
		_, err := db.Exec(migration)
		if err != nil {
			log.Fatal("Migration failed:", err)
		}
	}

	fmt.Println("Database migrations applied successfully.")
}
