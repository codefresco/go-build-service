package middleware

import (
	"time"

	jwt "github.com/codefresco/go-build-service/libs/token"
	"github.com/codefresco/go-build-service/loggerfactory"
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"go.uber.org/zap"
)

func RequestLogger() fiber.Handler {
	logger := loggerfactory.GetLogger()

	return func(c *fiber.Ctx) error {
		start := time.Now()

		err := c.Next()

		latency := time.Since(start).Milliseconds()

		logger.Info("request",
			zap.String("method", c.Method()),
			zap.String("path", c.Path()),
			zap.String("ip", c.IP()),
			zap.String("user_agent", c.Get("User-Agent")),
			zap.Int("status", c.Response().StatusCode()),
			zap.Int64("latency", latency),
			zap.String("request_id", c.Locals("requestid").(string)),
		)

		return err
	}
}

func Validator(body interface{}) fiber.Handler {
	return func(c *fiber.Ctx) error {
		validate := validator.New()

		if err := c.BodyParser(body); err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"error": "Invalid JSON",
			})
		}

		if err := validate.Struct(body); err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"error": err.Error(),
			})
		}

		c.Locals("body", body)
		return c.Next()
	}
}

func Authenticated() fiber.Handler {
	return func(c *fiber.Ctx) error {
		authHeader := c.Get("Authorization")

		claims, err := jwt.ValidateToken(authHeader)
		if err != nil {
			return c.Status(401).JSON(fiber.Map{
				"message": "Invalid token!",
				"error":   err.Error(),
			})
		}

		c.Locals("claims", claims)
		return c.Next()
	}
}
