package http

import (
	"context"

	"trade-union-service/internal/domains/users/domain/entity"

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

func NewHandler(service service, router fiber.Router, log *logrus.Logger) {
	handler := &Handler{
		service: service,
		log:     log,
	}

	handler.initRoutes(router)
}

func (h *Handler) initRoutes(router fiber.Router) {
	users := router.Group("/users")
	{
		users.Post("/new", h.createUser)
		users.Get("/user", h.getUser)
		users.Patch("/user", h.patchUser)
	}
}
