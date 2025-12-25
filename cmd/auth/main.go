package main

import (
	"auth/internal/config"
	"auth/pkg/logger"
)

func main() {
	cfg := config.MustLoad()

	log := logger.SetupLogger(cfg.Logger.Level)

	log.Info("cfg", cfg)
}