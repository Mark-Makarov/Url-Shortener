package main

import (
	"Url-Shortener/internal/config"
	"Url-Shortener/internal/http-server/handlers/redirect"
	"Url-Shortener/internal/http-server/handlers/url/save"
	"Url-Shortener/internal/http-server/middleware/loggerMiddleware"
	"Url-Shortener/internal/lib/logger/sl"
	"Url-Shortener/internal/lib/logger/slPretty"
	"net/http"

	"Url-Shortener/internal/storage/sqlite"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"log/slog"
	"os"
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

	router := chi.NewRouter()

	router.Use(middleware.RequestID)
	router.Use(middleware.RealIP)
	router.Use(loggerMiddleware.New(log))
	router.Use(middleware.Recoverer)
	router.Use(middleware.URLFormat)

	router.Route("/url", func(r chi.Router) {
		r.Use(middleware.BasicAuth("url-shortener", map[string]string{
			cfg.HTTPServer.User: cfg.HTTPServer.Password,
		}))

		r.Post("/", save.New(log, storage))
	})

	router.Get("/{alias}", redirect.New(log, storage))

	log.Info("server starting", slog.String("address", cfg.Address))

	server := &http.Server{
		Addr:        cfg.Address,
		Handler:     router,
		ReadTimeout: cfg.Timeout,
		IdleTimeout: cfg.IdleTimeout,
	}

	if err := server.ListenAndServe(); err != nil {
		log.Error("server start fail")
	}
}

func setupLogger(env string) *slog.Logger {
	var logger *slog.Logger

	switch env {
	case "development":
		logger = setupPrettySlog()
	case "production":
		logger = slog.New(slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelInfo}))
	}

	return logger
}

func setupPrettySlog() *slog.Logger {
	opts := slPretty.PrettyHandlerOptions{
		SlogOpts: &slog.HandlerOptions{
			Level: slog.LevelDebug,
		},
	}

	handler := opts.NewPrettyHandler(os.Stdout)

	return slog.New(handler)
}
