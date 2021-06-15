package server

import (
	"github.com/gofiber/fiber/v2"
	"github.com/joaovicdsantos/discord-clone-api/model"
	"github.com/joaovicdsantos/discord-clone-api/service"
	"github.com/joaovicdsantos/discord-clone-api/utils"
)

var (
	serverService service.ServerService
)

// GetAll get all servers
func GetAll(c *fiber.Ctx) error {
	servers := serverService.GetAll()
	return c.JSON(servers)
}

// GetOne get one specific server
func GetOne(c *fiber.Ctx) error {
	server, err := serverService.GetOne(c.Params("id"))
	if err.Err != nil {
		c.SendStatus(err.StatusCode)
		return c.JSON(fiber.Map{
			"error": err.Err.Error(),
		})
	}
	return c.JSON(server)
}

// GetAllChannelGroups find all channel groups by server id
func GetAllChannelGroups(c *fiber.Ctx) error {
	server, err := serverService.GetAllGroupChannels(c.Params("id"))
	if err.Err != nil {
		c.SendStatus(err.StatusCode)
		return c.JSON(fiber.Map{
			"error": err.Err.Error(),
		})
	}
	return c.JSON(server)

}

// Create create a new server
func Create(c *fiber.Ctx) error {
	var serverSave model.Server
	err := utils.ToModel(serverSave, c.BodyParser)
	if err.Err != nil {
		return c.Status(err.StatusCode).JSON(fiber.Map{
			"error": err.Err.Error(),
		})
	}
	server, err := serverService.Create(serverSave)
	if err.Err != nil {
		c.SendStatus(err.StatusCode)
		return c.JSON(fiber.Map{
			"error": err.Err.Error(),
		})
	}
	c.SendStatus(201)
	return c.JSON(server)
}

// Delete delete a server by id
func Delete(c *fiber.Ctx) error {
	err := serverService.Delete(c.Params("id"))
	if err.Err != nil {
		c.SendStatus(err.StatusCode)
		return c.JSON(fiber.Map{
			"error": err.Err.Error(),
		})
	}
	return c.SendStatus(204)
}

// Update update a server by id
func Update(c *fiber.Ctx) error {
	var serverUpdate model.Server
	err := utils.ToModel(serverUpdate, c.BodyParser)
	if err.Err != nil {
		return c.Status(err.StatusCode).JSON(fiber.Map{
			"error": err.Err.Error(),
		})
	}
	err = serverService.Update(c.Params("id"), serverUpdate)
	if err.Err != nil {
		c.SendStatus(err.StatusCode)
		return c.JSON(fiber.Map{
			"error": err.Err.Error(),
		})
	}
	return c.SendStatus(204)
}
