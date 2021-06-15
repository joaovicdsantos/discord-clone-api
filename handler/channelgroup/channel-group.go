package channelgroup

import (
	"github.com/gofiber/fiber/v2"
	"github.com/joaovicdsantos/discord-clone-api/model"
	"github.com/joaovicdsantos/discord-clone-api/service"
	"github.com/joaovicdsantos/discord-clone-api/utils"
)

var (
	groupChannelService service.ChannelGroupService
)

// GetAll get all channel groups
func GetAll(c *fiber.Ctx) error {
	channelGroups := groupChannelService.GetAll()
	return c.JSON(channelGroups)
}

// GetOne get one specific channel group
func GetOne(c *fiber.Ctx) error {
	channelGroup, err := groupChannelService.GetOne(c.Params("id"))
	if err.Err != nil {
		c.SendStatus(err.StatusCode)
		return c.JSON(fiber.Map{
			"error": err.Err.Error(),
		})
	}
	return c.JSON(channelGroup)
}

// Create create a new channel group
func Create(c *fiber.Ctx) error {
	var channelGroupSave model.ChannelGroup
	err := utils.ToModel(channelGroupSave, c.BodyParser)
	if err.Err != nil {
		return c.Status(err.StatusCode).JSON(fiber.Map{
			"error": err.Err.Error(),
		})
	}
	channelGroup, err := groupChannelService.Create(channelGroupSave)
	if err.Err != nil {
		c.SendStatus(err.StatusCode)
		return c.JSON(fiber.Map{
			"error": err.Err.Error(),
		})
	}
	c.SendStatus(201)
	return c.JSON(channelGroup)
}

// Delete delete a channel group by id
func Delete(c *fiber.Ctx) error {
	err := groupChannelService.Delete(c.Params("id"))
	if err.Err != nil {
		c.SendStatus(err.StatusCode)
		return c.JSON(fiber.Map{
			"error": err.Err.Error(),
		})
	}
	return c.SendStatus(204)
}

// Update update a channel group by id
func Update(c *fiber.Ctx) error {
	var channelGroupUpdate model.ChannelGroup
	err := utils.ToModel(channelGroupUpdate, c.BodyParser)
	if err.Err != nil {
		return c.Status(err.StatusCode).JSON(fiber.Map{
			"error": err.Err.Error(),
		})
	}
	err = groupChannelService.Update(c.Params("id"), channelGroupUpdate)
	if err.Err != nil {
		c.SendStatus(err.StatusCode)
		return c.JSON(fiber.Map{
			"error": err.Err.Error(),
		})
	}
	return c.SendStatus(204)
}
