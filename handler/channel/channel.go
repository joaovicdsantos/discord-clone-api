package channel

import (
	"github.com/gofiber/fiber/v2"
	"github.com/joaovicdsantos/discord-clone-api/model"
	"github.com/joaovicdsantos/discord-clone-api/service"
	"github.com/joaovicdsantos/discord-clone-api/utils"
)

var (
	channelService service.ChannelService
)

// GetAll get all channels
func GetAll(c *fiber.Ctx) error {
	channels := channelService.GetAll()
	return c.JSON(channels)
}

// GetOne get one specific channel
func GetOne(c *fiber.Ctx) error {
	channel, err := channelService.GetOne(c.Params("id"))
	if err.Err != nil {
		c.SendStatus(err.StatusCode)
		return c.JSON(fiber.Map{
			"error": err.Err.Error(),
		})
	}
	return c.JSON(channel)
}

// Create create a new channel
func Create(c *fiber.Ctx) error {
	var channelSave model.Channel
	err := utils.ToModel(channelSave, c.BodyParser)
	if err.Err != nil {
		return c.Status(err.StatusCode).JSON(fiber.Map{
			"error": err.Err.Error(),
		})
	}
	channel, err := channelService.Create(channelSave)
	if err.Err != nil {
		c.SendStatus(err.StatusCode)
		return c.JSON(fiber.Map{
			"error": err.Err.Error(),
		})
	}
	c.SendStatus(201)
	return c.JSON(channel)
}

// Delete delete a channel by id
func Delete(c *fiber.Ctx) error {
	err := channelService.Delete(c.Params("id"))
	if err.Err != nil {
		c.SendStatus(err.StatusCode)
		return c.JSON(fiber.Map{
			"error": err.Err.Error(),
		})
	}
	return c.SendStatus(204)
}

// Update update a channel by id
func Update(c *fiber.Ctx) error {
	var channelUpdate model.Channel
	err := utils.ToModel(channelUpdate, c.BodyParser)
	if err.Err != nil {
		return c.Status(err.StatusCode).JSON(fiber.Map{
			"error": err.Err.Error(),
		})
	}
	err = channelService.Update(c.Params("id"), channelUpdate)
	if err.Err != nil {
		c.SendStatus(err.StatusCode)
		return c.JSON(fiber.Map{
			"error": err.Err.Error(),
		})
	}
	return c.SendStatus(204)
}
