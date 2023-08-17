package main

import (
	"Url-Shortener/internal/config"
	"Url-Shortener/internal/http-server/middleware/loggerMiddleware"
	"Url-Shortener/internal/lib/logger"
	"Url-Shortener/internal/storage/sqlite"
	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"log/slog"
	"os"
)

func main() {
	cfg := config.MustLoad()
	log := setupLogger(cfg.Env)

	_, err := sqlite.New(cfg.StoragePath)
	if err != nil {
		log.Error("failed to init DB", sl.Err(err))
		os.Exit(1)
	}

	router := chi.NewRouter()

	router.Use(middleware.RequestID)
	router.Use(middleware.RealIP)
	router.Use(loggerMiddleware.New(log))
	router.Use(middleware.Recoverer)
	router.Use(middleware.URLFormat)
}

func setupLogger(env string) *slog.Logger {
	var logger *slog.Logger

	switch env {
	case "development":
		logger = slog.New(slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelDebug}))
	case "production":
		logger = slog.New(slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelInfo}))
	}

	return logger
}
