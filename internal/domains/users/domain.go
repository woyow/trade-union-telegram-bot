package users

import (
	"trade-union-service/internal/domains/users/delivery/http"
	"trade-union-service/internal/domains/users/domain/service"
	"trade-union-service/internal/domains/users/infra/mongodb"

	_ "trade-union-service/internal/domains/users/metrics"

	setupFiber "trade-union-service/internal/setup/fiber"
	setupMongo "trade-union-service/internal/setup/mongodb"

	"github.com/sirupsen/logrus"
)

const (
	domainLoggingKey = "domain"
	usersDomain      = "users"
)

func NewDomain(setupMongo *setupMongo.MongoDB, setupFiber *setupFiber.Fiber, log *logrus.Logger) {
	repo := mongodb.NewRepoImpl(setupMongo.Client, log)
	svc := service.NewService(repo, nil, log)
	http.InitHandler(svc, setupFiber.App, log)

	log.WithField(domainLoggingKey, usersDomain).
		Info("Domain has been initialized")
}
