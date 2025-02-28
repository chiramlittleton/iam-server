package auth

import (
	"database/sql"
	"net/http"

	"github.com/golang-jwt/jwt/v4"
)

// RBACMiddleware enforces role-based access control
func RBACMiddleware(requiredRole string, db *sql.DB) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			authHeader := r.Header.Get("Authorization")
			if authHeader == "" {
				http.Error(w, "Missing Authorization header", http.StatusUnauthorized)
				return
			}

			tokenString := authHeader[len("Bearer "):]
			claims := jwt.MapClaims{}
			token, err := jwt.ParseWithClaims(tokenString, &claims, func(token *jwt.Token) (interface{}, error) {
				return jwtSecret, nil
			})

			if err != nil || !token.Valid {
				http.Error(w, "Invalid token", http.StatusUnauthorized)
				return
			}

			email, ok := claims["email"].(string)
			if !ok {
				http.Error(w, "Invalid token claims", http.StatusUnauthorized)
				return
			}

			// Check user role in DB
			var role string
			err = db.QueryRow(
				"SELECT roles.name FROM user_roles JOIN roles ON user_roles.role_id = roles.id "+
					"JOIN users ON user_roles.user_id = users.id WHERE users.email=$1",
				email,
			).Scan(&role)

			if err != nil {
				http.Error(w, "Role not found", http.StatusForbidden)
				return
			}

			// Check if role matches
			if role != requiredRole {
				http.Error(w, "Access denied", http.StatusForbidden)
				return
			}

			next.ServeHTTP(w, r)
		})
	}
}
