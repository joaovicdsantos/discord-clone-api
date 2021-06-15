package user

import (
	"github.com/gofiber/fiber/v2"
	"github.com/joaovicdsantos/discord-clone-api/model"
	"github.com/joaovicdsantos/discord-clone-api/service"
	"github.com/joaovicdsantos/discord-clone-api/utils"
)

var (
	userService service.UserService
)

// GetAll get all users
func GetAll(c *fiber.Ctx) error {
	users := userService.GetAll()
	return c.JSON(users)
}

// GetOne get one specific user
func GetOne(c *fiber.Ctx) error {
	user, err := userService.GetOne(c.Params("id"))
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
	var userLogin model.User
	err := utils.ToModel(userLogin, c.BodyParser)
	if err.Err != nil {
		return c.Status(err.StatusCode).JSON(fiber.Map{
			"error": err.Err.Error(),
		})
	}
	token, err := userService.Login(userLogin)
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

	return nil
}

// Register register a new user
func Register(c *fiber.Ctx) error {
	var userRegister model.User
	err := utils.ToModel(userRegister, c.BodyParser)
	if err.Err != nil {
		return c.Status(err.StatusCode).JSON(fiber.Map{
			"error": err.Err.Error(),
		})
	}
	user, err := userService.Create(userRegister)
	if err.Err != nil {
		c.SendStatus(err.StatusCode)
		return c.JSON(fiber.Map{
			"error": err.Err.Error(),
		})
	}
	c.SendStatus(201)
	return c.JSON(user)
}

// Delete delete a user by id
func Delete(c *fiber.Ctx) error {
	err := userService.Delete(c.Params("id"))
	if err.Err != nil {
		c.SendStatus(err.StatusCode)
		return c.JSON(fiber.Map{
			"error": err.Err.Error(),
		})
	}
	return c.SendStatus(204)
}

// Update update a user by id
func Update(c *fiber.Ctx) error {
	var userUpdate model.User
	err := utils.ToModel(userUpdate, c.BodyParser)
	if err.Err != nil {
		return c.Status(err.StatusCode).JSON(fiber.Map{
			"error": err.Err.Error(),
		})
	}
	err = userService.Update(c.Params("id"), userUpdate)
	if err.Err != nil {
		c.SendStatus(err.StatusCode)
		return c.JSON(fiber.Map{
			"error": err.Err.Error(),
		})
	}
	return c.SendStatus(204)
}
