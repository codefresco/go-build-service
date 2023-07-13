package router

import (
	"time"

	"github.com/codefresco/go-build-service/api/auth"
	"github.com/gofiber/fiber/v2"
)

func Start(router *fiber.App) {

	router.Get("/", func(c *fiber.Ctx) error {
		response := map[string]interface{}{
			"message": "Works!",
			"status":  200,
		}
		time.Sleep(100 * time.Millisecond)
		return c.Status(200).JSON(response)
	})

	auth.UseAuthRouter(router.Group("/auth"))

	router.Use(func(c *fiber.Ctx) error {
		return c.Status(404).JSON(fiber.Map{
			"status":  404,
			"message": "404 Nothing to see here!",
		})
	})

}
