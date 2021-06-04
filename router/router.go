package router

import (
	"github.com/gofiber/fiber/v2"
	"github.com/joaovicdsantos/discord-clone-api/handler"
)

// SetupRoutes configure all routes
func SetupRoutes(app *fiber.App) {
	api := app.Group("/api/v1")

	serverRoutes := api.Group("/server")
	serverRoutes.Get("/", handler.GetServer)
	serverRoutes.Get("/:id", handler.GetServerById)
	serverRoutes.Post("/", handler.CreateServer)
	serverRoutes.Delete("/:id", handler.DeleteServer)
	serverRoutes.Put("/:id", handler.UpdateServer)

	channelRoutes := api.Group("/channel")
	channelRoutes.Get("/", handler.GetChannel)
	channelRoutes.Get("/:id", handler.GetChannelById)
	channelRoutes.Post("/", handler.CreateChannel)
	channelRoutes.Delete("/:id", handler.DeleteChannel)
	channelRoutes.Put("/:id", handler.UpdateChannel)

	channelGroupRoutes := api.Group("/channel-group")
	channelGroupRoutes.Get("/", handler.GetChannelGroup)
	channelGroupRoutes.Get("/:id", handler.GetChannelGroupById)
	channelGroupRoutes.Post("/", handler.CreateChannelGroup)
	channelGroupRoutes.Delete("/:id", handler.DeleteChannelGroup)
	channelGroupRoutes.Put("/:id", handler.UpdateChannelGroup)

}
