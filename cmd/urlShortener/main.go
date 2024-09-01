package main

import (
	slogg "URLshortener/internal/lib/logger/slog"
	"URLshortener/internal/parserConfig"
	"URLshortener/internal/storage/sqlite"
	"log"
	"log/slog"
	"os"
)

const (
	envLocal = "LOCAL"
	envDev   = "DEV"
	envProd  = "PROD"
)

func main() {
	configPath := "../../config/local.yaml"
	err := os.Setenv("CONFIG_PATH", configPath)
	if err != nil {
		log.Fatalf("failed to set CONFIG_PATH: %v", err)
	}

	cfg := parserConfig.MustLoad()
	logger := setupLogger(cfg.Env)
	logger.Info("start URL", slog.String("env", cfg.Env))
	logger.Debug("debug logger is on")

	storage, err := sqlite.New(cfg.StoragePath)
	if err != nil {
		logger.Error("failed to init storage", slogg.Err(err))
		os.Exit(1)
	}
	_ = storage
}

func setupLogger(env string) *slog.Logger {
	var logger *slog.Logger

	switch env {
	case envLocal:
		logger = slog.New(slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelDebug}))
	case envDev:
		logger = slog.New(slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelDebug}))
	case envProd:
		logger =
			slog.New(slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelInfo}))
	default:
		logger = slog.New(slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelInfo}))
	}

	return logger
}
