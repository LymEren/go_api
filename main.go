package main

import (
	"go_api/app"
	"go_api/configs"
	"go_api/repository"
	"go_api/services"

	"github.com/gofiber/fiber/v2"
)

func main() {

	appRoute := fiber.New()
	// Connect config file
	configs.ConnectDB()
	dbClient := configs.GetCollection(configs.DB, "todos")

	TodoRepositoryDB := repository.NewTodoRepositoryDB(dbClient)

	// Service for todoHandler

	td := app.TodoHandler{Service: services.NewTodoServise(TodoRepositoryDB)}

	// Post request (postman)
	appRoute.Post("/api/todos", td.CreateTodo)
	appRoute.Listen(":8080")
}
