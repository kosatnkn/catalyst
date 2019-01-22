package metrics

import "github.com/prometheus/client_golang/prometheus"

// Init registers declared metrics.
func Init() {
	prometheus.MustRegister(HTTPReqDuration)
}
