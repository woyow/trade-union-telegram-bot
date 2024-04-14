package victoria_metrics

import (
	"github.com/VictoriaMetrics/metrics"
	"github.com/sirupsen/logrus"
)

type VictoriaMetrics struct {
}

func NewVictoriaMetrics(log *logrus.Logger) {
	log.Println("victoria-metrics: ", metrics.ListMetricNames())
}
