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
	errGetUserAccessForbidden = errors.New("access forbidden")
)

func getUserAccess(c fiber.Ctx) (int, error) {
	token := c.Get("Authorization")

	if token != os.Getenv("ADMIN_API_TOKEN") {
		return http.StatusForbidden, errGetUserAccessForbidden
	}

	return 0, nil
}

type getUserRequest struct {
	ID     *string `form:"id,omitempty"`
	ChatID *int64  `form:"chat_id,omitempty"`
}

type getUserResponse struct {
	Result *entity.GetUserOut `json:"result"`
}

func (h *Handler) getUser(c fiber.Ctx) error {
	if statusCode, err := getUserAccess(c); err != nil {
		return errorResponse(c, statusCode, err)
	}

	var query getUserRequest

	if err := c.Bind().Query(&query); err != nil {
		return errorResponse(c, http.StatusBadRequest, err)
	}

	ctx, cancel := context.WithTimeout(c.Context(), 1*time.Second)
	defer cancel()

	out, err := h.service.GetUser(ctx, entity.GetUserServiceDTO{
		ID:     query.ID,
		ChatID: query.ChatID,
	})
	if err != nil {
		switch err {
		case errs.ErrUserNotFound:
			return errorResponse(c, http.StatusNotFound, err)
		default:
			return errorResponse(c, http.StatusInternalServerError, err)
		}
	}

	return c.Status(http.StatusOK).JSON(getUserResponse{
		Result: out,
	})
}
