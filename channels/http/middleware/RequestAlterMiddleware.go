package middleware

import (
	"context"
	"net/http"

	"github.com/google/uuid"

	"github.com/kosatnkn/catalyst/app/container"
	"github.com/kosatnkn/catalyst/domain/globals"
)

// RequestAlterMiddleware alerts the request.
type RequestAlterMiddleware struct {
	container *container.Container
}

// NewRequestAlterMiddleware returns a new instance of RequestAlterMiddleware.
func NewRequestAlterMiddleware(ctr *container.Container) *RequestAlterMiddleware {
	return &RequestAlterMiddleware{
		container: ctr,
	}
}

// Middleware executes middleware rules of RequestAlterMiddleware.
func (m *RequestAlterMiddleware) Middleware(next http.Handler) http.Handler {

	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		// create and attach a uuid to the request context
		contextUUID := uuid.New().String()

		ctx := context.WithValue(r.Context(), globals.UUIDKey, contextUUID)

		r = r.WithContext(ctx)

		// Call the next handler, which can be another middleware in the chain, or the final handler.
		next.ServeHTTP(w, r)
	})
}
