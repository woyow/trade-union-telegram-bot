package http

import (
	"github.com/gofiber/fiber/v3"
)

type errResponse struct {
	Error string `json:"error"`
}

func errorResponse(c fiber.Ctx, statusCode int, err error) error {
	return c.Status(statusCode).JSON(errResponse{Error: err.Error()})
}
