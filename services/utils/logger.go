package services_utils

import (
	"log/slog"
	"os"
	"sync"
)

// LoggerConfigurer is a struct that configures the logger for the server and uses slog
var (
	logger              *slog.Logger
	onceConfigureLogger sync.Once
)

type LoggerConfigurer struct {
	Logger   *slog.Logger
	LogLevel slog.Level
}

func NewLoggerConfigurer(logLevel slog.Level) LoggerConfigurer {
	onceConfigureLogger.Do(func() {
		logger = slog.New(slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{Level: logLevel}))
	})

	return LoggerConfigurer{
		Logger:   logger,
		LogLevel: logLevel,
	}
}

func GetLogger() *slog.Logger {
	return logger
}
