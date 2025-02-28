package api

import (
	"database/sql"

	"github.com/chiramlittleton/iam-server/api/handlers"
	"github.com/chiramlittleton/iam-server/internal/auth"
	"github.com/gorilla/mux"
)

// NewRouter initializes routes and applies middleware
func NewRouter(db *sql.DB) *mux.Router {
	router := mux.NewRouter()

	// Public routes (Multi-app support)
	router.HandleFunc("/register-client", handlers.RegisterUserForClient).Methods("POST")
	router.HandleFunc("/login-client", handlers.LoginUserForClient).Methods("POST")

	// Protected routes (JWT required)
	protected := router.PathPrefix("/api").Subrouter()
	protected.Use(auth.JWTMiddleware)
	protected.HandleFunc("/protected", handlers.ProtectedEndpoint).Methods("GET")

	// Admin routes (Requires "admin" role)
	admin := protected.PathPrefix("/admin").Subrouter()
	admin.Use(auth.RBACMiddleware("admin", db))
	admin.HandleFunc("/dashboard", handlers.AdminDashboard).Methods("GET")

	return router
}
