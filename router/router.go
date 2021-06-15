package router

import (
	"os"

	"github.com/gofiber/fiber/v2"
	jwtware "github.com/gofiber/jwt/v2"
	"github.com/joaovicdsantos/discord-clone-api/handler/channel"
	"github.com/joaovicdsantos/discord-clone-api/handler/channelgroup"
	"github.com/joaovicdsantos/discord-clone-api/handler/message"
	"github.com/joaovicdsantos/discord-clone-api/handler/server"
	"github.com/joaovicdsantos/discord-clone-api/handler/user"
)

// SetupRoutes configure all routes
func SetupRoutes(app *fiber.App) {

	api := app.Group("/api/v1")

	api.Post("/register", user.Register)
	api.Post("/login", user.Login)

	// JWT
	app.Use(jwtware.New(jwtware.Config{
		SigningKey:  []byte(os.Getenv("SECRET_KEY")),
		TokenLookup: "cookie:Authorization",
	}))

	// Server
	serverRoutes := api.Group("/server")
	serverRoutes.Get("/", server.GetAll)
	serverRoutes.Get("/:id", server.GetOne)
	serverRoutes.Get("/:id/channel-groups", server.GetAllChannelGroups)
	serverRoutes.Post("/", server.Create)
	serverRoutes.Delete("/:id", server.Delete)
	serverRoutes.Put("/:id", server.Update)

	// Channel
	channelRoutes := api.Group("/channel")
	channelRoutes.Get("/", channel.GetAll)
	channelRoutes.Get("/:id", channel.GetOne)
	channelRoutes.Post("/", channel.Create)
	channelRoutes.Delete("/:id", channel.Delete)
	channelRoutes.Put("/:id", channel.Update)

	// Channel Group
	channelGroupRoutes := api.Group("/channel-group")
	channelGroupRoutes.Get("/", channelgroup.GetAll)
	channelGroupRoutes.Get("/:id", channelgroup.GetOne)
	channelGroupRoutes.Post("/", channelgroup.Create)
	channelGroupRoutes.Delete("/:id", channelgroup.Delete)
	channelGroupRoutes.Put("/:id", channelgroup.Update)

	// User
	userRoutes := api.Group("/user")
	userRoutes.Get("/", user.GetAll)
	userRoutes.Get("/:id", user.GetOne)
	userRoutes.Delete("/:id", user.Delete)
	userRoutes.Put("/:id", user.Update)

	// Message
	messageRoutes := api.Group("/message")
	messageRoutes.Get("/", message.GetAll)
	messageRoutes.Get("/:id", message.GetOne)
	messageRoutes.Post("/", message.Create)
	messageRoutes.Delete("/:id", message.Delete)
	messageRoutes.Put("/:id", message.Update)
}
