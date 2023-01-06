package rest

import (
	"net/http"

	"github.com/gofiber/fiber/v2"
)

func (h handler) findAll(ctx *fiber.Ctx) error {
	todos, err := h.svc.FindAll()
	if err != nil {
		return err
	}
	return ctx.Status(http.StatusOK).JSON(todos)
}
