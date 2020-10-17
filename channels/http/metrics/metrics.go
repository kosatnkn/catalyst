package metrics

import "github.com/prometheus/client_golang/prometheus"

// Instrument and register all used metrics in this package.
// https://godoc.org/github.com/prometheus/client_golang/prometheus

// HTTPReqDuration instruments metrics for http request durations.
var HTTPReqDuration = prometheus.NewHistogramVec(
	prometheus.HistogramOpts{
		Name:    "http_request_duration_milliseconds",
		Help:    "HTTP request duration in milliseconds, partitioned by status code, HTTP method and URL.",
		Buckets: []float64{.5, 1, 2.5, 5, 10, 50, 100, 500, 1000},
	},
	[]string{"code", "method", "url"},
)
