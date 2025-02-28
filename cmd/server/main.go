package main

import (
	"log"
	"net/http"
	"os"

	"github.com/chiramlittleton/iam-server/api"
	"github.com/chiramlittleton/iam-server/internal/storage"
)

func main() {
	// Load port from environment variable
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	// Initialize database connection
	if err := storage.ConnectDB(); err != nil {
		log.Fatal(err)
	}

	// Initialize Router with database
	router := api.NewRouter(storage.DB) // ✅ Pass DB to `NewRouter()`

	// Start Server
	log.Printf("IAM Service running on port %s", port)
	err := http.ListenAndServe(":"+port, router)
	if err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}
