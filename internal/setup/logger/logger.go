package logger

import (
	"github.com/sirupsen/logrus"
)

func NewLogger(cfg *Config) *logrus.Logger {
	logger := logrus.StandardLogger()

	formatter := logrus.TextFormatter{
		DisableTimestamp: cfg.DisableTimestamp,
		FullTimestamp:    cfg.FullTimestamp,
	}

	logger.SetFormatter(&formatter)

	// Possible logLevel value: "panic", "fatal", "error", "warn" or "warning", "info", "debug", "trace"
	level, err := logrus.ParseLevel(cfg.Level)
	if err != nil {
		logger.WithError(err).WithField("level", cfg.Level).Warn("cannot parse a logging level")
	} else {
		logger.SetLevel(level)
	}

	return logger
}
