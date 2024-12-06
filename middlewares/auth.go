package middleware

import (
	"net/http"
	"strings"
	"stock_api/controllers"
	"fmt"
	"github.com/dgrijalva/jwt-go"
)

// Middleware to verify Token JWT
func AuthMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Get header Authorization token
		authHeader := r.Header.Get("Authorization")
		if authHeader == "" {
			http.Error(w, "Token not provided", http.StatusUnauthorized)
			return
		}

		// O token deve vir no formato: "Bearer <token>"
		tokenString := strings.TrimPrefix(authHeader, "Bearer ")

		// Parseia e valida o token
		token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
			// Verifica se o método de assinatura é correto (HS256)
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, fmt.Errorf("Invalid assinature method")
			}
			// Retorna a chave secreta para verificar o token
			return controllers.JwtSecret, nil
		})

		if err != nil || !token.Valid {
			http.Error(w, "Invalid token", http.StatusUnauthorized)
			return
		}

		next.ServeHTTP(w, r)
	})
}
