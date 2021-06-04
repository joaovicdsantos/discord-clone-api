package handler

import (
	"github.com/gofiber/fiber/v2"
	"github.com/joaovicdsantos/discord-clone-api/service"
)

var (
	groupChannelService service.ChannelGroupService
)

// GetChannelGroup get all group channels
func GetChannelGroup(c *fiber.Ctx) error {
	channelGroups := groupChannelService.FindAll()
	return c.JSON(channelGroups)
}

// GetChannelGroupById get one specific channelGroup
func GetChannelGroupById(c *fiber.Ctx) error {
	channelGroup, err := groupChannelService.FindById(c.Params("id"))
	if err.Err != nil {
		c.SendStatus(err.StatusCode)
		return c.JSON(fiber.Map{
			"error": err.Err.Error(),
		})
	}
	return c.JSON(channelGroup)
}

// CreateChannelGroup create a new channelGroup
func CreateChannelGroup(c *fiber.Ctx) error {
	channelGroup, err := groupChannelService.Create(c.BodyParser)
	if err.Err != nil {
		c.SendStatus(err.StatusCode)
		return c.JSON(fiber.Map{
			"error": err.Err.Error(),
		})
	}
	c.SendStatus(201)
	return c.JSON(channelGroup)
}

// DeleteChannelGroup delete a channelGroup by id
func DeleteChannelGroup(c *fiber.Ctx) error {
	err := groupChannelService.Delete(c.Params("id"))
	if err.Err != nil {
		c.SendStatus(err.StatusCode)
		return c.JSON(fiber.Map{
			"error": err.Err.Error(),
		})
	}
	return c.SendStatus(204)
}

// UpdateChannelGroup update a channelGroup by id
func UpdateChannelGroup(c *fiber.Ctx) error {
	err := groupChannelService.Update(c.Params("id"), c.BodyParser)
	if err.Err != nil {
		c.SendStatus(err.StatusCode)
		return c.JSON(fiber.Map{
			"error": err.Err.Error(),
		})
	}
	return c.SendStatus(204)
}
