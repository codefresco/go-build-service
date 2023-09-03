package user

import (
	"github.com/codefresco/go-build-service/middleware"
	"github.com/gofiber/fiber/v2"
)

func AuthRouter(router fiber.Router) {
	router.Post("/login", middleware.Validator(new(LoginUser)), Login)
	router.Post("/logout", middleware.Authenticated(), Logout)
	router.Post("/refresh", Refresh)
	router.Post("/register", middleware.Validator(new(RegisterUser)), Register)
}
