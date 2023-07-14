package user

import (
	"errors"

	"github.com/codefresco/go-build-service/libs/pass"
	"github.com/codefresco/go-build-service/loggerfactory"
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

func Register(c *fiber.Ctx) error {
	userDetails := c.Locals("body").(*RegisterUser)

	err := CreateUser(userDetails)

	if errors.Is(err, gorm.ErrDuplicatedKey) {
		return c.Status(409).JSON(fiber.Map{
			"message": "User exists!",
			"error":   err.Error(),
		})
	}

	if err != nil {
		return c.Status(500).JSON(fiber.Map{
			"message": "User registeration failed",
			"error":   err.Error(),
		})
	}

	return c.Status(200).JSON(fiber.Map{
		"message": "User registered!",
		"user":    userDetails,
	})
}

func Login(c *fiber.Ctx) error {
	logger := loggerfactory.GetSugaredLogger()
	loginDetails := c.Locals("body").(*LoginUser)

	dbUser, err := FindUser(loginDetails)
	if err != nil {
		return c.Status(404).JSON(fiber.Map{
			"message": "User not found!",
			"error":   err.Error(),
		})
	}

	validPassword, loginError := pass.PasswordChecks(loginDetails.Password, dbUser.PasswordHash, dbUser.PasswordSalt)
	if loginError != nil {
		logger.Errorw("Error checking password!", "error", loginError)
		return c.Status(500).JSON(fiber.Map{
			"message": "Permission check failed!",
			"error":   errors.New("Error checking permission!"),
		})
	}

	if !validPassword {
		return c.Status(403).JSON(fiber.Map{
			"message": "Permission denied!",
			"error":   errors.New("Access denied error!"),
		})
	}

	return c.Status(200).JSON(fiber.Map{
		"message": "Login successful!",
		"user":    loginDetails,
	})
}
