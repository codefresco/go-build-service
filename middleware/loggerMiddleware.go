package loggerMiddleware

import (
	"time"

	"github.com/codefresco/go-build-service/loggerFactory"
	"github.com/gofiber/fiber/v2"
	"go.uber.org/zap"
)

func RequestLogger() fiber.Handler {
	logger := loggerFactory.GetLogger()

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
