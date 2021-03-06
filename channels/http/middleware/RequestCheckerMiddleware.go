package middleware

import (
	"fmt"
	"net/http"

	"github.com/kosatnkn/catalyst/app/container"
	"github.com/kosatnkn/catalyst/channels/http/middleware/errors"
	"github.com/kosatnkn/catalyst/channels/http/response"
)

// RequestCheckerMiddleware validates the request header.
type RequestCheckerMiddleware struct {
	container     *container.Container
	omittedRoutes []string
}

// NewRequestCheckerMiddleware returns a new instance of RequestCheckerMiddleware.
func NewRequestCheckerMiddleware(ctr *container.Container) *RequestCheckerMiddleware {

	return &RequestCheckerMiddleware{
		container: ctr,
		omittedRoutes: []string{
			"/favicon.ico",
		},
	}
}

// Middleware executes middleware rules of RequestCheckerMiddleware.
func (m *RequestCheckerMiddleware) Middleware(next http.Handler) http.Handler {

	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		requestURI := r.RequestURI
		contentType := r.Header.Get("Content-Type")

		// skip omitted routes
		for _, v := range m.omittedRoutes {

			if v == requestURI {
				next.ServeHTTP(w, r)
				return
			}
		}

		// check content type
		if contentType != "application/json" {

			err := errors.NewMiddlewareError("100", fmt.Sprintf("API only accepts JSON as Content-Type, '%s' is given", contentType))

			response.Error(r.Context(), w, m.container.Adapters.Log, err)

			return
		}

		// Call the next handler, which can be another middleware in the chain, or the final handler.
		next.ServeHTTP(w, r)
	})
}
