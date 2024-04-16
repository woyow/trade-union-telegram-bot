package echotron

import (
	"github.com/NicoNex/echotron/v3"
	"github.com/sirupsen/logrus"
)

const (
	setupLoggingKey   = "setup"
	setupLoggingValue = "echotron"
)

type Echotron struct {
	API   *echotron.API
	Token string
}

func NewEchotron(cfg *Config, log *logrus.Logger) (*Echotron, error) {
	api := echotron.NewAPI(cfg.Token)

	log.WithField(setupLoggingKey, setupLoggingValue).
		Info("NewEchotron - API has been initialized")

	return &Echotron{
		API:   &api,
		Token: cfg.Token,
	}, nil
}
