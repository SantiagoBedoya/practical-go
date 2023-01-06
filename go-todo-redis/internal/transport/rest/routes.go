package rest

import (
	"github.com/SantiagoBedoya/todo-app/internal/models"
	"github.com/gofiber/fiber/v2"
)

type handler struct {
	svc models.TodoService
}

// RegisterRoutes add routes to server
func RegisterRoutes(app *fiber.App, svc models.TodoService) {
	apiV1 := app.Group("/api/v1/todos")
	{
		handler := handler{svc}
		apiV1.Get("/", handler.findAll)
		apiV1.Get("/:id", handler.findByID)
		apiV1.Post("/", handler.createTodo)
	}
}
