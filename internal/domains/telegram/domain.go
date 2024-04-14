package telegram

import (
	"trade-union-service/internal/domains/telegram/delivery/telegram"
	"trade-union-service/internal/domains/telegram/domain/service"
	"trade-union-service/internal/domains/telegram/infra/mongodb"
	"trade-union-service/internal/domains/telegram/infra/tgapi"

	_ "trade-union-service/internal/domains/telegram/metrics"

	setupEchotron "trade-union-service/internal/setup/echotron"
	setupMongo "trade-union-service/internal/setup/mongodb"

	"github.com/sirupsen/logrus"
)

const (
	domainLoggingKey = "domain"
	telegramDomain   = "telegram"
)

func NewDomain(setupMongo *setupMongo.MongoDB, setupEchotron *setupEchotron.Echotron, log *logrus.Logger) {
	repo := mongodb.NewRepoImpl(setupMongo.Client, log)
	api := tgapi.NewApiImpl(setupEchotron.API, log)
	svc := service.NewService(repo, api, log)
	bot := telegram.NewTelegram(svc, setupEchotron.Token, log)
	//metrics.InitMetrics()
	go func() {
		if err := bot.Run(); err != nil {
			log.WithField(domainLoggingKey, telegramDomain).
				Error("NewDomain - bot.Run error: ", err.Error())
			return
		}
	}()

	log.WithField(domainLoggingKey, telegramDomain).
		Info("Domain has been initialized")
}
