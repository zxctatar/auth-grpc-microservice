package logger

import (
	"log/slog"
	"os"
)

const (
	LevelDebug = "debug"
	LevelInfo = "info"
)

func SetupLogger(logLevel string) *slog.Logger {
	var log *slog.Logger

	switch logLevel {
	case LevelDebug:
		log = slog.New(slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelDebug}))
	case LevelInfo:
		log = slog.New(slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelInfo}))
	}

	return log
}