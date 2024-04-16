package tgapi

import (
	"github.com/NicoNex/echotron/v3"
	"github.com/sirupsen/logrus"
)

const (
	textLoggingKey     = "text"
	chatIDLoggingKey   = "chat_id"
	domainLoggingKey   = "domain"
	domainLoggingValue = "telegram"
	infraLoggingKey    = "infra"
	infraLoggingValue  = "tgapi"
)

type APIImpl struct {
	api *echotron.API
	log *logrus.Logger
}

func NewApiImpl(api *echotron.API, log *logrus.Logger) *APIImpl {
	return &APIImpl{
		api: api,
		log: log,
	}
}
