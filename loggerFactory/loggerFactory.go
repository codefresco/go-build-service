package loggerFactory

import (
	"log"

	"github.com/codefresco/go-build-service/config"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

func GetLogger() *zap.Logger {
	configs := config.GetConfig()

	var logDevelopment bool
	logOutputPaths := []string{"stdout"}
	logErrorPaths := []string{"stderr"}

	if configs.LogFilePath != "" {
		logOutputPaths = append(logOutputPaths, configs.LogFilePath)
		logErrorPaths = append(logErrorPaths, configs.LogFilePath)
	}

	if configs.Environment == "dev" {
		logDevelopment = true
	} else {
		logDevelopment = false
	}

	encoderConfig := zapcore.EncoderConfig{
		TimeKey:        "time",
		LevelKey:       "level",
		NameKey:        "logger",
		CallerKey:      "caller",
		MessageKey:     "message",
		StacktraceKey:  "stacktrace",
		LineEnding:     zapcore.DefaultLineEnding,
		EncodeLevel:    zapcore.LowercaseLevelEncoder,
		EncodeTime:     zapcore.ISO8601TimeEncoder,
		EncodeDuration: zapcore.SecondsDurationEncoder,
		EncodeCaller:   zapcore.ShortCallerEncoder,
	}

	config := zap.Config{
		Level:            zap.NewAtomicLevelAt(zap.InfoLevel),
		Development:      logDevelopment,
		Encoding:         "json",
		EncoderConfig:    encoderConfig,
		OutputPaths:      logOutputPaths,
		ErrorOutputPaths: logErrorPaths,
	}

	logger, err := config.Build(zap.AddCaller())
	if err != nil {
		log.Fatalf("could not initialize logger: %v", err)
		panic("Fatal error")
	}
	return logger
}
