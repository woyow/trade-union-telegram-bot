package telegram

import (
	"trade-union-service/internal/domains/telegram/delivery/telegram"
	"trade-union-service/internal/domains/telegram/domain/service"
	"trade-union-service/internal/domains/telegram/infra/mongodb"

	setupEchotron "trade-union-service/internal/setup/echotron"
	setupMongo "trade-union-service/internal/setup/mongodb"

	"github.com/sirupsen/logrus"
)

func NewDomain(setupMongo *setupMongo.MongoDB, setupEchotron *setupEchotron.Echotron, log *logrus.Logger) {
	repo := mongodb.NewRepoImpl(setupMongo.Client, log)
	svc := service.NewService(repo, setupEchotron.API, log)
	bot := telegram.NewBot(svc, setupEchotron.Token, log)

	go func() {
		if err := bot.Run(); err != nil {
			log.WithField("domain", "telegram").Error("NewDomain - bot.Run error: ", err.Error())
			return
		}
	}()
}