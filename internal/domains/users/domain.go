package users

import (
	"trade-union-service/internal/domains/users/delivery/http"
	"trade-union-service/internal/domains/users/domain/service"
	"trade-union-service/internal/domains/users/infra/mongodb"

	setupFiber "trade-union-service/internal/setup/fiber"
	setupMongo "trade-union-service/internal/setup/mongodb"

	"github.com/sirupsen/logrus"
)

func NewDomain(setupMongo *setupMongo.MongoDB, setupFiber *setupFiber.Fiber, log *logrus.Logger) {
	repo := mongodb.NewRepoImpl(setupMongo.Client, log)
	service := service.NewService(repo, repo, log)
	http.NewHandler(service, setupFiber.App, log)
}