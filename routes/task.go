package routes

import (
	"go-fiber-crud/handlers"

	"github.com/gofiber/fiber/v2"
)

func SetupTasksRoute(app *fiber.App) {
	tasksGroup := app.Group("/tasks")

	tasksGroup.Get("/", handlers.GetTasks)
	tasksGroup.Post("/", handlers.CreateTask)
	tasksGroup.Put("/:id", handlers.UpdateTask)
	tasksGroup.Get("/:id", handlers.GetTaskByID)
	tasksGroup.Delete("/:id", handlers.DeleteTask)
	tasksGroup.Patch("/:id/complete", handlers.CompleteTask)
}
