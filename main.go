package main

import (
	router "github.com/codefresco/go-build-service/api"
	"github.com/codefresco/go-build-service/config"
	"github.com/codefresco/go-build-service/loggerFactory"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	fiberLogger "github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/requestid"
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

	api.Use(requestid.New())
	api.Use(fiberLogger.New(fiberLogger.Config{
		Format: "${pid} ${locals:requestid} ${status} - ${method} ${path}​\n",
	}))

	logger.Info("Starting the api...")
	router.Start(api)
	api.Listen(":" + configs.Port)
}
