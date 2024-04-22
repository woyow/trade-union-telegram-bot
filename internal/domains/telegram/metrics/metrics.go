package metrics

import (
	setupVictoriaMetrics "trade-union-service/internal/setup/victoria-metrics"

	"github.com/VictoriaMetrics/metrics"
)

var metricsEnabled bool

func ConfigureMetrics(cfg *setupVictoriaMetrics.Config) {
	metricsEnabled = cfg.MetricsEnabled
}

var (
	messageTotalCounter              = metrics.NewCounter("tg_message_total")
	callbackQueryTotalCounter        = metrics.NewCounter("tg_callback_query_total")
	editedMessageTotalCounter        = metrics.NewCounter("tg_edited_message_total")
	businessErrorTotalCounter        = metrics.NewCounter("tg_business_error_total")
	errorTotalCounter                = metrics.NewCounter("tg_error_total")
	setStateTotalCounter             = metrics.NewCounter("tg_set_state_total")
	setStateErrorTotalCounter        = metrics.NewCounter("tg_set_state_error_total")
	setStateAndCallCounter           = metrics.NewCounter("tg_set_state_and_call_total")
	setStateAndCallErrorTotalCounter = metrics.NewCounter("tg_set_state_and_call_error_total")
)

func IncrementMessageTotal() {
	if metricsEnabled {
		messageTotalCounter.Inc()
	}
}

func IncrementCallbackQueryTotal() {
	if metricsEnabled {
		callbackQueryTotalCounter.Inc()
	}
}

func IncrementEditedMessageTotal() {
	if metricsEnabled {
		editedMessageTotalCounter.Inc()
	}
}

func IncrementBusinessErrorTotal() {
	if metricsEnabled {
		businessErrorTotalCounter.Inc()
	}
}

func IncrementErrorTotal() {
	if metricsEnabled {
		errorTotalCounter.Inc()
	}
}

func IncrementSetStateTotal() {
	if metricsEnabled {
		setStateTotalCounter.Inc()
	}
}

func IncrementSetStateErrorTotal() {
	if metricsEnabled {
		setStateErrorTotalCounter.Inc()
	}
}

func IncrementSetStateAndCallTotal() {
	if metricsEnabled {
		setStateAndCallCounter.Inc()
	}
}

func IncrementSetStateAndCallErrorTotal() {
	if metricsEnabled {
		setStateAndCallErrorTotalCounter.Inc()
	}
}
