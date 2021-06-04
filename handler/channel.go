package handler

import (
	"github.com/gofiber/fiber/v2"
	"github.com/joaovicdsantos/discord-clone-api/service"
)

var (
	channelService service.ChannelService
)

// GetChannel get all channels
func GetChannel(c *fiber.Ctx) error {
	channels := channelService.FindAll()
	return c.JSON(channels)
}

// GetChannelById get one specific channel
func GetChannelById(c *fiber.Ctx) error {
	channel, err := channelService.FindById(c.Params("id"))
	if err.Err != nil {
		c.SendStatus(err.StatusCode)
		return c.JSON(fiber.Map{
			"error": err.Err.Error(),
		})
	}
	return c.JSON(channel)
}

// CreateChannel create a new channel
func CreateChannel(c *fiber.Ctx) error {
	channel, err := channelService.Create(c.BodyParser)
	if err.Err != nil {
		c.SendStatus(err.StatusCode)
		return c.JSON(fiber.Map{
			"error": err.Err.Error(),
		})
	}
	c.SendStatus(201)
	return c.JSON(channel)
}

// DeleteChannel delete a channel by id
func DeleteChannel(c *fiber.Ctx) error {
	err := channelService.Delete(c.Params("id"))
	if err.Err != nil {
		c.SendStatus(err.StatusCode)
		return c.JSON(fiber.Map{
			"error": err.Err.Error(),
		})
	}
	return c.SendStatus(204)
}

// UpdateChannel update a channel by id
func UpdateChannel(c *fiber.Ctx) error {
	err := channelService.Update(c.Params("id"), c.BodyParser)
	if err.Err != nil {
		c.SendStatus(err.StatusCode)
		return c.JSON(fiber.Map{
			"error": err.Err.Error(),
		})
	}
	return c.SendStatus(204)
}
