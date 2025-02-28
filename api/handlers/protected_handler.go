package handlers

import (
	"encoding/json"
	"net/http"
)

// ProtectedEndpoint is an example of a JWT-secured route
func ProtectedEndpoint(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(map[string]string{"message": "You accessed a protected route!"})
}

