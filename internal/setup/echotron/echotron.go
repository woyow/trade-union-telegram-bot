package echotron

import (
	"github.com/NicoNex/echotron/v3"
	"github.com/sirupsen/logrus"
)

type Echotron struct {
	API   *echotron.API
	Token string
}

func NewEchotron(cfg *Config, log *logrus.Logger) (*Echotron, error) {
	api := echotron.NewAPI(cfg.Token)

	log.Info("setup echotron: NewEchotron - API has been initialized")

	return &Echotron{
		API:   &api,
		Token: cfg.Token,
	}, nil
}
