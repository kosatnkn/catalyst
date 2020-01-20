package metrics

import (
	httpMetrics "github.com/kosatnkn/catalyst/app/http/metrics"
	"github.com/prometheus/client_golang/prometheus"
)

// registerMetrics registers declared metrics.
//
// This is the central location to register metrics from different
// layers of the service.
func registerMetrics() {
	prometheus.MustRegister(httpMetrics.HTTPReqDuration)
}
