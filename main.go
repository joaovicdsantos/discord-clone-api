package main

import (
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/joaovicdsantos/discord-clone-api/database"
	"github.com/joaovicdsantos/discord-clone-api/router"
	"github.com/joho/godotenv"
)

func main() {

	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal(err)
	}

	app := fiber.New()

	// Log
	app.Use(logger.New())

	// CORS
	app.Use(cors.New(cors.Config{
		AllowCredentials: true,
	}))
	// Routes
	router.SetupRoutes(app)

	// InitDatabase
	database.InitDatabase()

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello, World!")
	})

	log.Fatal(app.Listen("localhost:3000"))
}
