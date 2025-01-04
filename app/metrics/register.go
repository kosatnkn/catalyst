package metrics

import (
	"github.com/kosatnkn/catalyst/v3/app/transport/http/middleware"
	"github.com/prometheus/client_golang/prometheus"
)

// Register registers declared custom metrics.
//
// This is the central location to register metrics from different
// layers of the service.
func Register() {
	// http_request_duration_milliseconds metric recorded by the MetricsMiddleware
	prometheus.MustRegister(middleware.HTTPReqDuration)
}
