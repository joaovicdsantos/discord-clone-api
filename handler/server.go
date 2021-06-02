package handler

import (
	"github.com/gofiber/fiber/v2"
	"github.com/joaovicdsantos/discord-clone-api/service"
	"github.com/joaovicdsantos/discord-clone-api/util"
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
	id, convertErr := util.VerifyAndConvertID(c.Params("id"))
	if convertErr.Err != nil {
		c.SendStatus(convertErr.StatusCode)
		return c.JSON(fiber.Map{
			"error": convertErr.Err.Error(),
		})
	}
	server, err := serverService.FindById(id)
	if err != nil {
		c.SendStatus(404)
		return c.JSON(fiber.Map{
			"error": err.Error(),
		})
	}
	return c.JSON(server)
}

// CreateServer create a new server
func CreateServer(c *fiber.Ctx) error {
	server, err := serverService.Create(c.BodyParser)
	if err != nil {
		c.SendStatus(404)
		return c.JSON(fiber.Map{
			"error": err.Error(),
		})
	}
	c.SendStatus(201)
	return c.JSON(server)
}

// DeleteServer delete a server by id
func DeleteServer(c *fiber.Ctx) error {
	id, convertErr := util.VerifyAndConvertID(c.Params("id"))
	if convertErr.Err != nil {
		c.SendStatus(convertErr.StatusCode)
		return c.JSON(fiber.Map{
			"error": convertErr.Err.Error(),
		})
	}
	serverService.Delete(uint(id))
	return c.SendStatus(204)
}

// UpdateServer update a server by id
func UpdateServer(c *fiber.Ctx) error {
	id, convertErr := util.VerifyAndConvertID(c.Params("id"))
	if convertErr.Err != nil {
		c.SendStatus(convertErr.StatusCode)
		return c.JSON(fiber.Map{
			"error": convertErr.Err.Error(),
		})
	}
	err := serverService.Update(id, c.BodyParser)
	if err != nil {
		c.SendStatus(400)
		return c.JSON(fiber.Map{
			"error": err.Error(),
		})
	}
	return c.SendStatus(204)
}
