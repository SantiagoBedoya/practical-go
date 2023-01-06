package rest

import (
	"github.com/SantiagoBedoya/todo-app/internal/models"
	"github.com/gofiber/fiber/v2"
)

func (h handler) createTodo(ctx *fiber.Ctx) error {
	todo := models.Todo{}
	if err := ctx.BodyParser(&todo); err != nil {
		return err
	}
	if err := h.svc.Create(todo); err != nil {
		return err
	}
	return ctx.JSON("todo created successfully")
}
