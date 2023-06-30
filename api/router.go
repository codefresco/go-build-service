package router

import (
	"github.com/gofiber/fiber/v2"
)

func Start(router *fiber.App) {

	router.Get("/", func(c *fiber.Ctx) error {
		return c.Status(200).SendString("Works!")
	})

	router.Use(func(c *fiber.Ctx) error {
		return c.Status(404).JSON(fiber.Map{
			"code":    404,
			"message": "404 Nothing to see here!",
		})
	})

}
