package handler

import (
	"github.com/gofiber/fiber/v2"
	"github.com/joaovicdsantos/discord-clone-api/service"
)

var (
	serverService service.ServerService
)

// GetServer get all servers
func GetServer(c *fiber.Ctx) error {
	servers := serverService.FindAll()
	return c.JSON(servers)
}

// GetServerById get one specific server
func GetServerById(c *fiber.Ctx) error {
	server, err := serverService.FindById(c.Params("id"))
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
	server, err := serverService.FindAllGroupChannels(c.Params("id"))
	if err.Err != nil {
		c.SendStatus(err.StatusCode)
		return c.JSON(fiber.Map{
			"error": err.Err.Error(),
		})
	}
	return c.JSON(server)

}

// CreateServer create a new server
func CreateServer(c *fiber.Ctx) error {
	server, err := serverService.Create(c.BodyParser)
	if err.Err != nil {
		c.SendStatus(err.StatusCode)
		return c.JSON(fiber.Map{
			"error": err.Err.Error(),
		})
	}
	c.SendStatus(201)
	return c.JSON(server)
}

// DeleteServer delete a server by id
func DeleteServer(c *fiber.Ctx) error {
	err := serverService.Delete(c.Params("id"))
	if err.Err != nil {
		c.SendStatus(err.StatusCode)
		return c.JSON(fiber.Map{
			"error": err.Err.Error(),
		})
	}
	return c.SendStatus(204)
}

// UpdateServer update a server by id
func UpdateServer(c *fiber.Ctx) error {
	err := serverService.Update(c.Params("id"), c.BodyParser)
	if err.Err != nil {
		c.SendStatus(err.StatusCode)
		return c.JSON(fiber.Map{
			"error": err.Err.Error(),
		})
	}
	return c.SendStatus(204)
}
