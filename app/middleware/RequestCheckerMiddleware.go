package middleware

import (
	"fmt"
	"net/http"

	"github.com/kosatnkn/catalyst/app/container"
	errTypes "github.com/kosatnkn/catalyst/app/error/types"
	"github.com/kosatnkn/catalyst/app/transport/response"
)

// RequestCheckerMiddleware validates the request header.
type RequestCheckerMiddleware struct {
	container     *container.Container
	omittedRoutes []string
}

// Init initialize a new instance of RequestCheckerMiddleware.
func (m *RequestCheckerMiddleware) Init(ctr *container.Container) {

	m.container = ctr

	m.omittedRoutes = []string{
		"/favicon.ico",
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

			err := errTypes.NewMiddlewareError(fmt.Sprintf("API only accepts JSON as Content-Type, '%s' is given", contentType), 100, "")

			response.Error(r.Context(), w, err, m.container.Adapters.Log)

			return
		}

		// Call the next handler, which can be another middleware in the chain, or the final handler.
		next.ServeHTTP(w, r)
	})
}
