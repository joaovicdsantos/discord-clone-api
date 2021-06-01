package router

import (
	"github.com/gofiber/fiber/v2"
	"github.com/joaovicdsantos/discord-clone-api/handler"
)

func SetupRoutes(app *fiber.App) {
	api := app.Group("/api/v1")

	serverRoutes := api.Group("/server")
	serverRoutes.Get("/", handler.GetServer)
	serverRoutes.Get("/:id", handler.GetServerById)
	serverRoutes.Post("/", handler.CreateServer)
	serverRoutes.Delete("/:id", handler.DeleteServer)
	serverRoutes.Put("/:id", handler.UpdateServer)
}
