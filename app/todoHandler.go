package app

import (
	"net/http"

	"github.com/goapiexamples/goeren/models"
	"github.com/goapiexamples/goeren/services"
	"github.com/gofiber/fiber/v2"
)

type TodoHandler struct {
	Service services.TodoService
}

func (h TodoHandler) CreateTodo(c *fiber.Ctx) error {

	var todo models.Todo

	// JSON data type --> todo struct
	if err := c.BodyParser(&todo); err != nil {
		return c.Status(http.StatusBadRequest).JSON(err.Error())
	}

	result, err := h.Service.TodoInsert(todo)

	if err != nil || result.Status == false {
		return err
	}

	return c.Status(http.StatusCreated).JSON(result)
}
