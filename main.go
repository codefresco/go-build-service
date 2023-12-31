package main

import (
	router "github.com/codefresco/go-build-service/api"
	"github.com/codefresco/go-build-service/config"
	postgres "github.com/codefresco/go-build-service/database"
	"github.com/codefresco/go-build-service/loggerfactory"
	"github.com/codefresco/go-build-service/middleware"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/helmet"
	"github.com/gofiber/fiber/v2/middleware/recover"
	"github.com/gofiber/fiber/v2/middleware/requestid"
	"github.com/joho/godotenv"
)

func main() {
	godotenv.Load()
	configs := config.GetConfig()
	logger := loggerfactory.GetSugaredLogger()

	postgres.Connect()

	api := fiber.New()
	api.Use(helmet.New())
	api.Use(recover.New())

	api.Use(cors.New(cors.Config{
		AllowOrigins: configs.AllowOrigins,
		AllowHeaders: configs.AllowHeaders,
	}))

	api.Use(requestid.New())
	api.Use(middleware.RequestLogger())

	logger.Infow("Starting:", "port", configs.Port)
	router.Start(api)
	api.Listen(":" + configs.Port)
}
