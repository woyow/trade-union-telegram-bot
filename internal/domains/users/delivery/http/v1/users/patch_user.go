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

var (
	errPatchUserAccessForbidden = errors.New("patch user access forbidden")
	errChatIDMinValidation      = errors.New("chat id validation error, min=1")
)

func patchUserAccess(c fiber.Ctx) (int, error) {
	token := c.Get("Authorization")

	if token != os.Getenv("ADMIN_API_TOKEN") {
		return http.StatusForbidden, errPatchUserAccessForbidden
	}

	return 0, nil
}

func (r *patchUserRequest) validate() (int, error) {
	for i := range r.Roles {
		switch r.Roles[i] {
		case adminRole, juristRole:
			continue
		default:
			return http.StatusBadRequest, errInvalidRole
		}
	}

	if *r.ChatID < 1 {
		return http.StatusBadRequest, errChatIDMinValidation
	}

	return 0, nil
}

type patchUserRequest struct {
	ID       string   `json:"id" validate:"required"`
	Roles    []string `json:"roles,omitempty"`
	Fname    *string  `json:"fname,omitempty"`
	Lname    *string  `json:"lname,omitempty"`
	Mname    *string  `json:"mname,omitempty"`
	Position *string  `json:"position,omitempty"`
	ChatID   *int64   `json:"chat_id,omitempty"`
}

type patchUserResponse struct {
	Result *entity.GetUserOut `json:"result"`
}

func (h *Handler) patchUser(c fiber.Ctx) error {
	if statusCode, err := patchUserAccess(c); err != nil {
		return errorResponse(c, statusCode, err)
	}

	var query patchUserRequest

	if err := c.Bind().JSON(&query); err != nil {
		return errorResponse(c, http.StatusBadRequest, err)
	}

	if statusCode, err := query.validate(); err != nil {
		return errorResponse(c, statusCode, err)
	}

	ctx, cancel := context.WithTimeout(c.Context(), 1*time.Second)
	defer cancel()

	out, err := h.service.UpdateUser(ctx, entity.UpdateUserServiceDTO{
		ID:       query.ID,
		Roles:    query.Roles,
		Fname:    query.Fname,
		Lname:    query.Lname,
		Mname:    query.Mname,
		Position: query.Position,
		ChatID:   query.ChatID,
	})
	if err != nil {
		switch err {
		case errs.ErrFieldRequiredForUpdate:
			return errorResponse(c, http.StatusBadRequest, err)
		case errs.ErrUserWithChatIDAlreadyExists:
			return errorResponse(c, http.StatusForbidden, err)
		default:
			return errorResponse(c, http.StatusInternalServerError, err)
		}
	}

	return c.Status(http.StatusCreated).JSON(patchUserResponse{
		Result: out,
	})
}
