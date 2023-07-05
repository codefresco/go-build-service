package loggerFactory

import (
	"log"
	"os"
	"sync"

	"github.com/codefresco/go-build-service/config"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

var (
	onceLogger    sync.Once
	logger        *zap.Logger
	onceSugared   sync.Once
	sugaredLogger *zap.SugaredLogger
)

func CreateLogger() *zap.Logger {
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
		CallerKey:      "source",
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
		InitialFields: map[string]interface{}{
			"pid":     os.Getpid(),
			"version": configs.Version,
		},
	}

	loggerInstance, err := config.Build(zap.AddCaller())
	if err != nil {
		log.Fatalf("could not initialize logger: %v", err)
		panic("Fatal error")
	}
	return loggerInstance
}

func CreateSugaredLogger() *zap.SugaredLogger {
	return CreateLogger().Sugar()
}

func GetLogger() *zap.Logger {
	onceLogger.Do(func() {
		logger = CreateLogger()
	})
	return logger
}

func GetSugaredLogger() *zap.SugaredLogger {
	onceSugared.Do(func() {
		sugaredLogger = CreateSugaredLogger()
	})
	return sugaredLogger
}
