package middleware

import (
	"context"
	"net/http"

	"github.com/google/uuid"

	"github.com/kosatnkn/catalyst/domain/globals"
)

// RequestAlterMiddleware alrers the request.
type RequestAlterMiddleware struct{}

// Middleware executes middleware rules of RequestAlterMiddleware.
func (rtm *RequestAlterMiddleware) Middleware(next http.Handler) http.Handler {

	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		// create and attach a uuid to the request context
		contextUUID := uuid.New().String()

		ctx := context.WithValue(r.Context(), globals.UUIDKey, contextUUID)

		r = r.WithContext(ctx)

		// Call the next handler, which can be another middleware in the chain, or the final handler.
		next.ServeHTTP(w, r)
	})
}
