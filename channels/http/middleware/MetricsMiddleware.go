package middleware

import (
	"net/http"
	"strconv"
	"strings"
	"time"

	"github.com/kosatnkn/catalyst/v2/channels/http/metrics"
)

// MetricsMiddleware attaches metrics to the request.
type MetricsMiddleware struct{}

// NewMetricsMiddleware returns a new instance of MetricsMiddleware.
func NewMetricsMiddleware() *MetricsMiddleware {
	return &MetricsMiddleware{}
}

// Middleware executes middleware rules of MetricsMiddleware.
func (m *MetricsMiddleware) Middleware(next http.Handler) http.Handler {

	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		startTime := time.Now()
		lrw := newLoggingResponseWriter(w)

		// Call the next handler, which can be another middleware in the chain, or the final handler.
		next.ServeHTTP(lrw, r)

		duration := float64(time.Since(startTime).Nanoseconds() / 1000000)
		obs := metrics.HTTPReqDuration.WithLabelValues(strconv.Itoa(lrw.statusCode), r.Method, m.generalizePath(r.URL.Path))
		obs.Observe(duration)
	})
}

// generalizePath generates a common signature for the given route endpoint.
//
// This is to avoid creating a large number of redundant metric values using path variables.
// Such metrics will carry little value and will have to be aggregated afterwords. And it will
// add unwanted amount of useless metric data.
// By generalizing the route endpoint metrics will be properly aggregated.
//
// '/resource/123', '/resource/456' and '/resource/-789' will be converted to '/resource/id'
// '/resource/79.5' and '/resource/-5.5' will be converted to '/resource/val'
// '/resource/123/lon/79.5/lat/5.5' will be converted to '/resource/id/lon/val/lat/val'
func (m *MetricsMiddleware) generalizePath(path string) string {

	routeParts := strings.Split(path, "/")

	for i, routePart := range routeParts {

		_, errInt := strconv.ParseInt(routePart, 10, 64)
		if errInt == nil {
			routeParts[i] = "id"
			continue
		}

		_, errFloat := strconv.ParseFloat(routePart, 64)
		if errFloat == nil {
			routeParts[i] = "val"
			continue
		}
	}

	return strings.Join(routeParts, "/")
}

// The loggingResponseWriter is created embedding http.ResponseWriter
// https://golang.org/doc/effective_go.html#embedding
// https://ndersson.me/post/capturing_status_code_in_net_http/
type loggingResponseWriter struct {
	http.ResponseWriter
	statusCode int
}

// newLoggingResponseWriter creates a new instance of loggingResponseWriter.
func newLoggingResponseWriter(w http.ResponseWriter) *loggingResponseWriter {
	return &loggingResponseWriter{w, http.StatusOK}
}

// WriteHeader intercepts the actual WriteHeader function of ResponseWriter
// and stores the status code in statusCode.
func (lrw *loggingResponseWriter) WriteHeader(code int) {
	lrw.statusCode = code
	lrw.ResponseWriter.WriteHeader(code)
}
