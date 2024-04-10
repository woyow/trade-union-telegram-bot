package http

import (
	"context"
	"errors"
	"net/http"
	"os"
	"time"

	"trade-union-service/internal/domains/users/domain/entity"
	"trade-union-service/internal/domains/users/errs"

	"github.com/gofiber/fiber/v3"
)

const (
	adminRole  = "admin"
	juristRole = "jurist"
)

var (
	errCreateUserAccessForbidden = errors.New("access forbidden")
	errInvalidRole               = errors.New("invalid role")
)

func createUserAccess(c fiber.Ctx) (int, error) {
	token := c.Get("Authorization")

	if token != os.Getenv("ADMIN_API_TOKEN") {
		return http.StatusForbidden, errCreateUserAccessForbidden
	}

	return 0, nil
}

func (r *createUserRequest) validate() (int, error) {
	for i := range r.Roles {
		switch r.Roles[i] {
		case adminRole, juristRole:
			continue
		default:
			return http.StatusBadRequest, errInvalidRole
		}
	}
	return 0, nil
}

type createUserRequest struct {
	Roles    []string `json:"roles"`
	Fname    string   `json:"fname"`
	Lname    string   `json:"lname"`
	Mname    string   `json:"mname"`
	Position string   `json:"position"`
	ChatID   int64    `json:"chat_id"`
}

type createUserResponse struct {
	Result *entity.CreateUserOut `json:"result"`
}

func (h *Handler) createUser(c fiber.Ctx) error {
	if statusCode, err := createUserAccess(c); err != nil {
		return errorResponse(c, statusCode, err)
	}

	var query createUserRequest

	if err := c.Bind().JSON(&query); err != nil {
		return errorResponse(c, http.StatusBadRequest, err)
	}

	if statusCode, err := query.validate(); err != nil {
		return errorResponse(c, statusCode, err)
	}

	ctx, cancel := context.WithTimeout(c.Context(), 1*time.Second)
	defer cancel()

	out, err := h.service.CreateUser(ctx, entity.CreateUserServiceDTO{
		Roles:    query.Roles,
		Fname:    query.Fname,
		Lname:    query.Lname,
		Mname:    query.Mname,
		Position: query.Position,
		ChatID:   query.ChatID,
	})
	if err != nil {
		switch err {
		case errs.ErrUserWithChatIDAlreadyExists:
			return errorResponse(c, http.StatusForbidden, err)
		default:
			return errorResponse(c, http.StatusInternalServerError, err)
		}
	}

	return c.Status(http.StatusCreated).JSON(createUserResponse{
		Result: out,
	})
}
