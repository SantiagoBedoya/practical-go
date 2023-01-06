package rest

import (
	"net/http"

	"github.com/gofiber/fiber/v2"
)

func (h handler) findByID(ctx *fiber.Ctx) error {
	todoID := ctx.Params("id")
	result, err := h.svc.FindByID(todoID)
	if err != nil {
		return err
	}
	if result == nil {
		return ctx.SendStatus(http.StatusNotFound)
	}
	return ctx.JSON(result)
}
