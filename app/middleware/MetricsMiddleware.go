package middleware

import (
	"net/http"
	"strconv"
	"time"

	"github.com/kosatnkn/catalyst/app/metrics"
)

// MetricsMiddleware alrers the request.
type MetricsMiddleware struct{}

// Middleware executes middleware rules of MetricsMiddleware.
func (rtm *MetricsMiddleware) Middleware(next http.Handler) http.Handler {

	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		startTime := time.Now()
		lrw := newLoggingResponseWriter(w)

		// Call the next handler, which can be another middleware in the chain, or the final handler.
		next.ServeHTTP(lrw, r)

		duration := float64(time.Since(startTime).Nanoseconds() / 1000000)
		metrics.HTTPReqDuration.WithLabelValues(strconv.Itoa(lrw.statusCode), r.Method, r.URL.Path).Observe(duration)
	})
}

// The loggingResponseWriter is created embedding http.ResponseWriter
// https://golang.org/doc/effective_go.html#embedding
// https://ndersson.me/post/capturing_status_code_in_net_http/
type loggingResponseWriter struct {
	http.ResponseWriter
	statusCode int
}

func newLoggingResponseWriter(w http.ResponseWriter) *loggingResponseWriter {
	return &loggingResponseWriter{w, http.StatusOK}
}

func (lrw *loggingResponseWriter) WriteHeader(code int) {
	lrw.statusCode = code
	lrw.ResponseWriter.WriteHeader(code)
}
