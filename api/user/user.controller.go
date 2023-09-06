package user

import (
	jwt "github.com/codefresco/go-build-service/libs/token"
	"github.com/gofiber/fiber/v2"
)

func GetUserDetails(c *fiber.Ctx) error {
	claims := c.Locals("claims").(jwt.MapClaims)

	userDetails, err := FindUser(&UserCredentials{Email: claims["sub"].(string), Password: ""})
	if err != nil {
		return userErrorHandler(c, err)
	}

	return c.Status(200).JSON(fiber.Map{
		"message":       "User details retreived!",
		"user_email":    userDetails.Email,
		"user_fullname": userDetails.FirstName + userDetails.LastName,
	})
}

func userErrorHandler(c *fiber.Ctx, err error) error {
	switch err {
	case ErrNotFound:
		return c.Status(404).JSON(fiber.Map{
			"message": "User does not exist!",
			"error":   err.Error(),
		})
	case jwt.ErrUnauthorized:
		return c.Status(403).JSON(fiber.Map{
			"message": "Invalid token!",
			"error":   err.Error(),
		})
	default:
		return c.Status(500).JSON(fiber.Map{
			"message": "Something went wrong!",
			"error":   ErrInternal.Error(),
		})
	}
}
