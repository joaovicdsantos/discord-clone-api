package main

import (
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/joaovicdsantos/discord-clone-api/database"
	"github.com/joho/godotenv"
)

func main() {

	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal(err)
	}

	database.InitDatabase()

	app := fiber.New()

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello, World!")
	})

	log.Fatal(app.Listen("localhost:3000"))
}
