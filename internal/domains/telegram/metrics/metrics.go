package metrics

import (
	"github.com/VictoriaMetrics/metrics"
)

var (
	MessageTotalCounter = metrics.NewCounter("tg_message_total")
	MessageErrorCounter = metrics.NewCounter("tg_message_error_total")
)

func MessageTotalIncrement() {
	MessageTotalCounter.Inc()
}

func MessageErrorIncrement() {
	MessageErrorCounter.Inc()
}
