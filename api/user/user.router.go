package user

import (
	"github.com/codefresco/go-build-service/middleware"
	"github.com/gofiber/fiber/v2"
)

func UserRouter(router fiber.Router) {
	router.Get("/details", middleware.Authenticated(), GetUserDetails)
}
