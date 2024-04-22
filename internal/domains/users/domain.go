package users

import (
	"trade-union-service/internal/domains/telegram/metrics"
	"trade-union-service/internal/domains/users/delivery/http"
	"trade-union-service/internal/domains/users/domain/service"
	"trade-union-service/internal/domains/users/infra/mongodb"

	setupFiber "trade-union-service/internal/setup/fiber"
	setupMongo "trade-union-service/internal/setup/mongodb"
	setupVictoriaMetrics "trade-union-service/internal/setup/victoria-metrics"

	"github.com/sirupsen/logrus"
)

const (
	domainLoggingKey   = "domain"
	domainLoggingValue = "users"
)

func NewDomain(setupMongo *setupMongo.MongoDB, setupFiber *setupFiber.Fiber, setupVictoriaMetrics *setupVictoriaMetrics.VictoriaMetrics, log *logrus.Logger) {
	metrics.ConfigureMetrics(setupVictoriaMetrics.Config)

	repo := mongodb.NewRepoImpl(setupMongo.Client, log)
	svc := service.NewService(repo, nil, log)
	http.InitHandler(svc, setupFiber.App, log)

	log.WithField(domainLoggingKey, domainLoggingValue).
		Info("Domain has been initialized")
}
