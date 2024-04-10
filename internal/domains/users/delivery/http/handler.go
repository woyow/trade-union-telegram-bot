package http

import (
	"context"
	"github.com/sirupsen/logrus"

	"trade-union-service/internal/domains/users/domain/entity"

	"github.com/gofiber/fiber/v3"
)

type service interface {
	CreateUser(ctx context.Context, dto entity.CreateUserServiceDTO) (*entity.CreateUserOut, error)
	GetUser(ctx context.Context, dto entity.GetUserServiceDTO) (*entity.GetUserOut, error)
}

type Handler struct {
	service service
	log     *logrus.Logger
}

func NewHandler(service service, app *fiber.App, log *logrus.Logger) {
	handler := &Handler{
		service: service,
		log:     log,
	}

	handler.initRoutes(app)
}

func (h *Handler) initRoutes (app *fiber.App) {
	v1 := app.Group("/v1")
	{
		users := v1.Group("/users")
		{
			users.Post("/new", h.createUser)
			users.Get("/user", h.getUser)
		}
	}
}
