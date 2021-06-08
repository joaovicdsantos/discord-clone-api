package router

import (
	"os"

	"github.com/gofiber/fiber/v2"
	jwtware "github.com/gofiber/jwt/v2"
	"github.com/joaovicdsantos/discord-clone-api/handler"
)

// SetupRoutes configure all routes
func SetupRoutes(app *fiber.App) {

	api := app.Group("/api/v1")

	api.Post("/register", handler.CreateUser)
	api.Post("/login", handler.Login)

	// JWT
	app.Use(jwtware.New(jwtware.Config{
		SigningKey:  []byte(os.Getenv("SECRET_KEY")),
		TokenLookup: "cookie:Authorization",
	}))

	// Server
	serverRoutes := api.Group("/server")
	serverRoutes.Get("/", handler.GetServer)
	serverRoutes.Get("/:id", handler.GetServerById)
	serverRoutes.Post("/", handler.CreateServer)
	serverRoutes.Delete("/:id", handler.DeleteServer)
	serverRoutes.Put("/:id", handler.UpdateServer)

	// Channel
	channelRoutes := api.Group("/channel")
	channelRoutes.Get("/", handler.GetChannel)
	channelRoutes.Get("/:id", handler.GetChannelById)
	channelRoutes.Post("/", handler.CreateChannel)
	channelRoutes.Delete("/:id", handler.DeleteChannel)
	channelRoutes.Put("/:id", handler.UpdateChannel)

	// Channel Group
	channelGroupRoutes := api.Group("/channel-group")
	channelGroupRoutes.Get("/", handler.GetChannelGroup)
	channelGroupRoutes.Get("/:id", handler.GetChannelGroupById)
	channelGroupRoutes.Post("/", handler.CreateChannelGroup)
	channelGroupRoutes.Delete("/:id", handler.DeleteChannelGroup)
	channelGroupRoutes.Put("/:id", handler.UpdateChannelGroup)

	// User
	userRoutes := api.Group("/user")
	userRoutes.Get("/", handler.GetUser)
	userRoutes.Get("/:id", handler.GetUserById)
	userRoutes.Delete("/:id", handler.DeleteUser)
	userRoutes.Put("/:id", handler.UpdateUser)

	// Message
	messageRoutes := api.Group("/message")
	messageRoutes.Get("/", handler.GetMessage)
	messageRoutes.Get("/:id", handler.GetMessageById)
	messageRoutes.Post("/", handler.CreateMessage)
	messageRoutes.Delete("/:id", handler.DeleteMessage)
	messageRoutes.Put("/:id", handler.UpdateMessage)
}
