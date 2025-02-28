package handlers

import (
	"encoding/json"
	"net/http"
)

// AdminDashboard handles requests to the admin dashboard
func AdminDashboard(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(map[string]string{"message": "Welcome to the admin dashboard"})
}
