package handlers

import (
	"context"
	"net/http"

	"go-web-test/internal/logger"
)

type contextKey string

const userContextKey contextKey = "user"

// AuthMiddleware
func AuthMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		logger.JSONLogger("info", "AuthMiddleware called")
		ctx := context.WithValue(r.Context(), userContextKey, "mek")
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}
