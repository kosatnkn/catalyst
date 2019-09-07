package middleware

import (
	"fmt"
	"net/http"

	"github.com/kosatnkn/catalyst/app/container"
	"github.com/kosatnkn/catalyst/app/error"
	errTypes "github.com/kosatnkn/catalyst/app/error/types"
)

// RequestCheckerMiddleware validates the request header.
type RequestCheckerMiddleware struct {
	container     *container.Container
	omittedRoutes []string
}

// Init initialize a new instance of RequestCheckerMiddleware.
func (rtm *RequestCheckerMiddleware) Init(ctr *container.Container) {

	rtm.container = ctr

	rtm.omittedRoutes = []string{
		"/favicon.ico",
	}
}

// Middleware executes middleware rules of RequestCheckerMiddleware.
func (rtm *RequestCheckerMiddleware) Middleware(next http.Handler) http.Handler {

	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		requestURI := r.RequestURI
		contentType := r.Header.Get("Content-Type")

		// skip omitted routes
		for _, v := range rtm.omittedRoutes {

			if v == requestURI {
				next.ServeHTTP(w, r)
				return
			}
		}

		// check content type
		if contentType != "application/json" {

			error.Handle(r.Context(), errTypes.NewMiddlewareError(fmt.Sprintf("API only accepts JSON, '%s' is given", contentType), 100, ""), w, rtm.container.Adapters.Log)

			return
		}

		// Call the next handler, which can be another middleware in the chain, or the final handler.
		next.ServeHTTP(w, r)
	})
}
