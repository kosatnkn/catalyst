package middleware

import (
	"net/http"
)

// CORSMiddleware attaches necessary CORS headers to the request.
type CORSMiddleware struct{}

// NewCORSMiddleware returns a new instance of CORSMiddleware.
func NewCORSMiddleware() *CORSMiddleware {
	return &CORSMiddleware{}
}

// Middleware executes middleware rules of CORSMiddleware.
func (m *CORSMiddleware) Middleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Add("Access-Control-Allow-Headers", "Content-Type,AccessToken,X-CSRF-Token, Authorization, Token")
		w.Header().Add("Access-Control-Allow-Credentials", "true")
		w.Header().Set("content-type", "application/json;charset=UTF-8")

		if r.Method == http.MethodOptions {
			w.WriteHeader(http.StatusNoContent)
			return
		}

		next.ServeHTTP(w, r)
	})
}
