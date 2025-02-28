package main

import (
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
)

func main() {
	// Load port from environment variable
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	// Initialize Router
	router := mux.NewRouter()

	// Health Check
	router.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("IAM Service is Running"))
	}).Methods("GET")

	// Start Server
	log.Printf("IAM Service running on port %s", port)
	err := http.ListenAndServe(":"+port, router)
	if err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}
