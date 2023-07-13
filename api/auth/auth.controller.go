package auth

import (
	"github.com/codefresco/go-build-service/loggerfactory"
	"github.com/gofiber/fiber/v2"
)

func Register(c *fiber.Ctx) error {
	logger := loggerfactory.GetSugaredLogger()
	user := c.Locals("body").(*RegisterUser)

	logger.Infof("info received %s %s, %s and %s", user.FirstName, user.LastName, user.Email, user.Password)

	return c.Status(200).JSON(fiber.Map{
		"message": "User registered!",
		"user":    user,
	})
}

func Login(c *fiber.Ctx) error {
	logger := loggerfactory.GetSugaredLogger()
	user := c.Locals("body").(*LoginUser)

	logger.Infof("info received %s and %s", user.Email, user.Password)

	return c.Status(200).JSON(fiber.Map{
		"message": "Login successful!",
		"user":    user,
	})
}
