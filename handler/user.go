package handler

import (
	"github.com/gofiber/fiber/v2"
	"github.com/joaovicdsantos/discord-clone-api/service"
)

var (
	userService service.UserService
)

// GetUser get all users
func GetUser(c *fiber.Ctx) error {
	users := userService.FindAll()
	return c.JSON(users)
}

// GetUserById get one specific user
func GetUserById(c *fiber.Ctx) error {
	user, err := userService.FindById(c.Params("id"))
	if err.Err != nil {
		c.SendStatus(err.StatusCode)
		return c.JSON(fiber.Map{
			"error": err.Err.Error(),
		})
	}
	return c.JSON(user)
}

// Login with a user
func Login(c *fiber.Ctx) error {
	token, err := userService.Login(c.BodyParser)
	if err.Err != nil {
		return c.Status(err.StatusCode).JSON(fiber.Map{
			"error": err.Err.Error(),
		})
	}

	cookie := new(fiber.Cookie)
	cookie.Name = "Authorization"
	cookie.Value = token
	cookie.HTTPOnly = true
	cookie.SameSite = "Strict"
	c.Cookie(cookie)

	return c.SendStatus(200)
}

// CreateUser create a new user
func CreateUser(c *fiber.Ctx) error {
	user, err := userService.Create(c.BodyParser)
	if err.Err != nil {
		c.SendStatus(err.StatusCode)
		return c.JSON(fiber.Map{
			"error": err.Err.Error(),
		})
	}
	c.SendStatus(201)
	return c.JSON(user)
}

// DeleteUser delete a user by id
func DeleteUser(c *fiber.Ctx) error {
	err := userService.Delete(c.Params("id"))
	if err.Err != nil {
		c.SendStatus(err.StatusCode)
		return c.JSON(fiber.Map{
			"error": err.Err.Error(),
		})
	}
	return c.SendStatus(204)
}

// UpdateUser update a user by id
func UpdateUser(c *fiber.Ctx) error {
	err := userService.Update(c.Params("id"), c.BodyParser)
	if err.Err != nil {
		c.SendStatus(err.StatusCode)
		return c.JSON(fiber.Map{
			"error": err.Err.Error(),
		})
	}
	return c.SendStatus(204)
}
