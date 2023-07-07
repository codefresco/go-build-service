package postgres

import (
	"time"

	"github.com/codefresco/go-build-service/config"
	"github.com/codefresco/go-build-service/loggerFactory"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func Connect() {
	var err error

	logger := loggerFactory.GetSugaredLogger()
	configs := config.GetConfig()

	logger.Infow("Connecting to postgres...")

	for attempt := 1; attempt <= 3; attempt++ {
		DB, err = gorm.Open(postgres.Open(configs.PostgresUrl), &gorm.Config{})
		if err != nil {
			logger.Warnw("Postgres connection attempt failed. Retrying...", "attempt", attempt)
			time.Sleep(time.Second * time.Duration(attempt*attempt))

			continue
		}
	}

	if err != nil {
		logger.Fatalw("Failed to connected to postgres! Terminating...", "error", err)
	}
}
