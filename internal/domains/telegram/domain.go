package telegram

import (
	"trade-union-service/internal/domains/telegram/delivery/telegram"
	"trade-union-service/internal/domains/telegram/domain/service"
	"trade-union-service/internal/domains/telegram/infra/mongodb"
	"trade-union-service/internal/domains/telegram/infra/tgapi"
	"trade-union-service/internal/domains/telegram/metrics"
	setupEchotron "trade-union-service/internal/setup/echotron"
	setupMongo "trade-union-service/internal/setup/mongodb"
	setupVictoriaMetrics "trade-union-service/internal/setup/victoria-metrics"

	"github.com/sirupsen/logrus"
)

const (
	domainLoggingKey   = "domain"
	domainLoggingValue = "telegram"
)

func NewDomain(setupMongo *setupMongo.MongoDB, setupEchotron *setupEchotron.Echotron, setupVictoriaMetrics *setupVictoriaMetrics.VictoriaMetrics, log *logrus.Logger) {
	metrics.ConfigureMetrics(setupVictoriaMetrics.Config)

	repo := mongodb.NewRepoImpl(setupMongo.Client, log)
	api := tgapi.NewApiImpl(setupEchotron.API, log)
	svc := service.NewService(repo, api, log)
	bot := telegram.NewTelegram(svc, setupEchotron.Token, log)

	go func() {
		if err := bot.Run(); err != nil {
			log.WithField(domainLoggingKey, domainLoggingValue).
				Error("NewDomain - bot.Run error: ", err.Error())

			return
		}
	}()

	log.WithField(domainLoggingKey, domainLoggingValue).
		Info("Domain has been initialized")
}
