package metrics

import (
	setupVictoriaMetrics "trade-union-service/internal/setup/victoria-metrics"
)

var metricsEnabled bool

func ConfigureMetrics(cfg *setupVictoriaMetrics.Config) {
	metricsEnabled = cfg.MetricsEnabled
}
