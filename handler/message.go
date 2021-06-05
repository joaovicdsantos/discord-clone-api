package handler

import (
	"github.com/gofiber/fiber/v2"
	"github.com/joaovicdsantos/discord-clone-api/service"
)

var (
	messageService service.MessageService
)

// GetMessage get all messages
func GetMessage(c *fiber.Ctx) error {
	messages := messageService.FindAll()
	return c.JSON(messages)
}

// GetMessageById get one specific message
func GetMessageById(c *fiber.Ctx) error {
	message, err := messageService.FindById(c.Params("id"))
	if err.Err != nil {
		c.SendStatus(err.StatusCode)
		return c.JSON(fiber.Map{
			"error": err.Err.Error(),
		})
	}
	return c.JSON(message)
}

// CreateMessage create a new message
func CreateMessage(c *fiber.Ctx) error {
	message, err := messageService.Create(c.BodyParser)
	if err.Err != nil {
		c.SendStatus(err.StatusCode)
		return c.JSON(fiber.Map{
			"error": err.Err.Error(),
		})
	}
	c.SendStatus(201)
	return c.JSON(message)
}

// DeleteMessage delete a message by id
func DeleteMessage(c *fiber.Ctx) error {
	err := messageService.Delete(c.Params("id"))
	if err.Err != nil {
		c.SendStatus(err.StatusCode)
		return c.JSON(fiber.Map{
			"error": err.Err.Error(),
		})
	}
	return c.SendStatus(204)
}

// UpdateMessage update a message by id
func UpdateMessage(c *fiber.Ctx) error {
	err := messageService.Update(c.Params("id"), c.BodyParser)
	if err.Err != nil {
		c.SendStatus(err.StatusCode)
		return c.JSON(fiber.Map{
			"error": err.Err.Error(),
		})
	}
	return c.SendStatus(204)
}
