package main

import (
	"context"
	"go-fiber-crud/database"
	"go-fiber-crud/routes"

	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New()

	database.InitDB()
	defer func() {
		if err := database.Client.Disconnect(context.TODO()); err != nil {
			panic(err)
		}
	}()

	routes.SetupTasksRoute(app)

	app.Listen(":3333")
}
