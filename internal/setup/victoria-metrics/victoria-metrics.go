package victoria_metrics

import (
	"fmt"
	"github.com/VictoriaMetrics/metrics"
	"github.com/sirupsen/logrus"
	"time"
)

const (
	setupLoggingKey   = "setup"
	setupLoggingValue = "victoria-metrics"
	proto             = "http"
	separator         = ":"
)

type VictoriaMetrics struct {
	Config *Config
	log    *logrus.Logger
}

func NewVictoriaMetrics(cfg *Config, log *logrus.Logger) *VictoriaMetrics {
	log.WithField(setupLoggingKey, setupLoggingValue).
		Debug("Initialized metrics: ", metrics.ListMetricNames())

	return &VictoriaMetrics{
		Config: cfg,
		log:    log,
	}
}

func (m *VictoriaMetrics) Run() error {
	if m.Config.MetricsEnabled {
		pushURL := fmt.Sprintf("%s://%s/api/v1/import/prometheus", proto, m.Config.Host+separator+m.Config.Port)

		if err := metrics.InitPush(
			pushURL,
			time.Duration(m.Config.PushInterval)*time.Second,
			`instance="trade-union"`,
			true,
		); err != nil {
			m.log.WithField(setupLoggingKey, setupLoggingValue).
				Error("Run - metrics.InitPush error: ", err.Error())

			return err
		}

		m.log.WithField(setupLoggingKey, setupLoggingValue).
			Info("Run - metrics has been initialized")
	} else {
		m.log.WithField(setupLoggingKey, setupLoggingValue).
			Info("Run - metrics disabled")
	}

	return nil
}
