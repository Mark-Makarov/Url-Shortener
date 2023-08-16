package main

import (
	"Url-Shortener/internal/config"
	"Url-Shortener/internal/lib/logger"
	"Url-Shortener/internal/storage/sqlite"
	"log/slog"
	"os"
)

const (
	envDev  = "development"
	envProd = "production"
)

func main() {
	cfg := config.MustLoad()
	log := setupLogger(cfg.Env)

	storage, err := sqlite.New(cfg.StoragePath)
	if err != nil {
		log.Error("failed to init DB", sl.Err(err))
		os.Exit(1)
	}

	_ = storage
}

func setupLogger(env string) *slog.Logger {
	var logger *slog.Logger

	switch env {
	case envDev:
		logger = slog.New(slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelDebug}))
	case envProd:
		logger = slog.New(slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelInfo}))
	}

	return logger
}
