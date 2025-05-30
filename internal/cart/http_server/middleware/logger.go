package middleware

import (
	"cart/internal/cart/logger"
	"net/http"
)

func Logger(next http.Handler) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		logger.Infof(r.Context(), "start to handle request method: %s, url: %s", r.Method, r.URL.Path)
		defer logger.Infof(r.Context(), "finished handle request method: %s, url: %s", r.Method, r.URL.Path)
		next.ServeHTTP(w, r)
	}
}
