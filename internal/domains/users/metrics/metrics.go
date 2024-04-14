package metrics

import (
	"github.com/VictoriaMetrics/metrics"
)

var (
	TestMetric = metrics.NewCounter("test")
)
