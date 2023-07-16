package postgres

import (
	"time"

	"github.com/codefresco/go-build-service/config"
	"github.com/codefresco/go-build-service/loggerfactory"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func Connect() {
	var err error

	logger := loggerfactory.GetSugaredLogger()
	configs := config.GetConfig()

	for attempt := 0; attempt < 3; attempt++ {
		time.Sleep(time.Second * time.Duration(attempt*attempt))
		logger.Infow("Connecting to postgres...", "attempt", attempt+1)

		DB, err = gorm.Open(postgres.Open(configs.PostgresURL), &gorm.Config{})
		if err == nil {
			logger.Infow("Connected to postgres!")
			return
		}
		logger.Warnw("Failed to connected to postgres!", "error", err)
	}

	logger.Fatalw("Failed to connected to postgres! Terminating...", "error", err)
}
