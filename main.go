package main

import (
	"github.com/goapiexamples/goeren/app"
	"github.com/goapiexamples/goeren/configs"
	"github.com/goapiexamples/goeren/repository"
	"github.com/goapiexamples/goeren/services"
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
	appRoute.Post("/api/todo", td.CreateTodo)
	appRoute.Listen(":8080")
}
