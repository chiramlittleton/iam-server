package handlers

import (
	"database/sql"
	"encoding/json"
	"net/http"

	"github.com/chiramlittleton/iam-server/internal/auth"
	"github.com/chiramlittleton/iam-server/internal/storage"
)

type RegisterRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
	ClientID int    `json:"client_id"` // Required for multi-application support
}

func RegisterUserForClient(w http.ResponseWriter, r *http.Request) {
	var req RegisterRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Invalid request", http.StatusBadRequest)
		return
	}

	hashedPassword, err := auth.HashPassword(req.Password)
	if err != nil {
		http.Error(w, "Error hashing password", http.StatusInternalServerError)
		return
	}

	// Insert the new user or fetch existing user ID
	var userID int
	err = storage.DB.QueryRow(
		"INSERT INTO users (email, password_hash) VALUES ($1, $2) ON CONFLICT (email) DO UPDATE SET email=EXCLUDED.email RETURNING id",
		req.Email, hashedPassword,
	).Scan(&userID)
	if err != nil {
		http.Error(w, "Error creating user", http.StatusInternalServerError)
		return
	}

	// Link the user to the requested application
	_, err = storage.DB.Exec(
		"INSERT INTO user_profiles (user_id, client_id, role) VALUES ($1, $2, 'user') ON CONFLICT DO NOTHING",
		userID, req.ClientID,
	)
	if err != nil {
		http.Error(w, "Error assigning user profile", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(map[string]string{"message": "User registered for client"})
}

func LoginUserForClient(w http.ResponseWriter, r *http.Request) {
	var req RegisterRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Invalid request", http.StatusBadRequest)
		return
	}

	var userID int
	var hashedPassword string
	err := storage.DB.QueryRow(
		"SELECT users.id, users.password_hash FROM users "+
			"JOIN user_profiles ON users.id = user_profiles.user_id "+
			"WHERE users.email=$1 AND user_profiles.client_id=$2",
		req.Email, req.ClientID,
	).Scan(&userID, &hashedPassword)

	if err == sql.ErrNoRows {
		http.Error(w, "Invalid credentials or not registered for this client", http.StatusUnauthorized)
		return
	} else if err != nil {
		http.Error(w, "Database error", http.StatusInternalServerError)
		return
	}

	if !auth.CheckPassword(hashedPassword, req.Password) {
		http.Error(w, "Invalid credentials", http.StatusUnauthorized)
		return
	}

	token, err := auth.GenerateJWT(userID, req.Email, req.ClientID)
	if err != nil {
		http.Error(w, "Error generating token", http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(map[string]string{"token": token})
}
