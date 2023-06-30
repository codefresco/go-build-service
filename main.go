package main

import (
	router "github.com/codefresco/go-build-service/api"
	"github.com/codefresco/go-build-service/config"
	"github.com/codefresco/go-build-service/loggerFactory"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/joho/godotenv"
)

func main() {
	godotenv.Load()
	configs := config.GetConfig()
	logger := loggerFactory.GetLogger()

	api := fiber.New()
	api.Use(cors.New(cors.Config{
		AllowOrigins: configs.AllowOrigins,
		AllowHeaders: configs.AllowHeaders,
	}))

	logger.Info("Starting the api...")
	router.Start(api)
	api.Listen(":" + configs.Port)
}
