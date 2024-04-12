package http

import (
	"context"
	"trade-union-service/internal/domains/users/domain/entity"

	usersV1 "trade-union-service/internal/domains/users/delivery/http/v1/users"

	"github.com/gofiber/fiber/v3"
	"github.com/sirupsen/logrus"
)

type service interface {
	CreateUser(ctx context.Context, dto entity.CreateUserServiceDTO) (*entity.CreateUserOut, error)
	GetUser(ctx context.Context, dto entity.GetUserServiceDTO) (*entity.GetUserOut, error)
	UpdateUser(ctx context.Context, dto entity.UpdateUserServiceDTO) (*entity.GetUserOut, error)
}

type Handler struct {
	service service
	log     *logrus.Logger
}

func InitHandler(service service, app *fiber.App, log *logrus.Logger) {
	handler := &Handler{
		service: service,
		log:     log,
	}

	handler.initRoutes(service, app, log)
}

func (h *Handler) initRoutes(service service, app *fiber.App, log *logrus.Logger) {
	v1 := app.Group("/v1")

	usersV1.NewHandler(service, v1, log)
}
