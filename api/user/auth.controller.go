package user

import (
	"github.com/codefresco/go-build-service/libs/pass"
	"github.com/codefresco/go-build-service/loggerfactory"
	"github.com/gofiber/fiber/v2"
)

func Register(c *fiber.Ctx) error {
	userDetails := c.Locals("body").(*RegisterUser)

	err := CreateUser(userDetails)
	if err != nil {
		return authErrorHandler(c, err)
	}

	return c.Status(200).JSON(fiber.Map{
		"message": "User registered!",
		"user":    userDetails.Email,
	})
}

func Login(c *fiber.Ctx) error {
	logger := loggerfactory.GetSugaredLogger()
	loginDetails := c.Locals("body").(*LoginUser)

	dbUser, err := FindUser(loginDetails)
	if err != nil {
		return authErrorHandler(c, err)
	}

	validPassword, loginError := pass.PasswordChecks(loginDetails.Password, dbUser.PasswordHash, dbUser.PasswordSalt)
	if loginError != nil {
		logger.Errorw("Error checking password!", "error", loginError)
		return authErrorHandler(c, ErrInternal)
	}

	if !validPassword {
		return authErrorHandler(c, ErrPermissionDenied)
	}

	return c.Status(200).JSON(fiber.Map{
		"message": "Login successful!",
		"user":    loginDetails.Email,
	})
}

func authErrorHandler(c *fiber.Ctx, err error) error {
	switch err {
	case ErrAlreadyEsists:
		return c.Status(409).JSON(fiber.Map{
			"message": "Could not create user!",
			"error":   err.Error(),
		})
	case ErrNotFound:
		return c.Status(404).JSON(fiber.Map{
			"message": "User does not exist!",
			"error":   err.Error(),
		})
	case ErrPermissionDenied:
		return c.Status(403).JSON(fiber.Map{
			"message": "Invalid authentication details!",
			"error":   err.Error(),
		})
	default:
		return c.Status(500).JSON(fiber.Map{
			"message": "Something went wrong!",
			"error":   ErrInternal.Error(),
		})
	}
}
