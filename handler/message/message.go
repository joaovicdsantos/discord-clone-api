package message

import (
	"github.com/gofiber/fiber/v2"
	"github.com/joaovicdsantos/discord-clone-api/model"
	"github.com/joaovicdsantos/discord-clone-api/service"
	"github.com/joaovicdsantos/discord-clone-api/utils"
)

var (
	messageService service.MessageService
)

// GetAll get all messages
func GetAll(c *fiber.Ctx) error {
	messages := messageService.GetAll()
	return c.JSON(messages)
}

// GetOne get one specific message
func GetOne(c *fiber.Ctx) error {
	message, err := messageService.GetOne(c.Params("id"))
	if err.Err != nil {
		c.SendStatus(err.StatusCode)
		return c.JSON(fiber.Map{
			"error": err.Err.Error(),
		})
	}
	return c.JSON(message)
}

// Create create a new message
func Create(c *fiber.Ctx) error {
	var messageSave model.Message
	err := utils.ToModel(messageSave, c.BodyParser)
	if err.Err != nil {
		return c.Status(err.StatusCode).JSON(fiber.Map{
			"error": err.Err.Error(),
		})
	}
	message, err := messageService.Create(messageSave)
	if err.Err != nil {
		c.SendStatus(err.StatusCode)
		return c.JSON(fiber.Map{
			"error": err.Err.Error(),
		})
	}
	c.SendStatus(201)
	return c.JSON(message)
}

// Delete delete a message by id
func Delete(c *fiber.Ctx) error {
	err := messageService.Delete(c.Params("id"))
	if err.Err != nil {
		c.SendStatus(err.StatusCode)
		return c.JSON(fiber.Map{
			"error": err.Err.Error(),
		})
	}
	return c.SendStatus(204)
}

// Update update a message by id
func Update(c *fiber.Ctx) error {
	var messageUpdate model.Message
	err := utils.ToModel(messageUpdate, c.BodyParser)
	if err.Err != nil {
		return c.Status(err.StatusCode).JSON(fiber.Map{
			"error": err.Err.Error(),
		})
	}
	err = messageService.Update(c.Params("id"), messageUpdate)
	if err.Err != nil {
		c.SendStatus(err.StatusCode)
		return c.JSON(fiber.Map{
			"error": err.Err.Error(),
		})
	}
	return c.SendStatus(204)
}
