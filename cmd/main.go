package main

import (
	"context"
	"fmt"
	"log/slog"
	"net"
	"os"

	"github.com/ta8i2chi8/go-api-sample/internal/config"
	"github.com/ta8i2chi8/go-api-sample/internal/presentation/router"
	"github.com/ta8i2chi8/go-api-sample/internal/server"
	"github.com/ta8i2chi8/go-api-sample/pkg/logger"
)

func main() {
	logger.Init(slog.LevelDebug)

	slog.Info("starting api server...")
	if err := run(context.Background()); err != nil {
		slog.Error("error occurred while running the server", "err", err)
		os.Exit(1)
	}
}

func run(ctx context.Context) error {
	cfg, err := config.Load(ctx)
	if err != nil {
		slog.Error("failed to load config", "err", err)
		return err
	}

	l, err := net.Listen("tcp", fmt.Sprintf(":%s", cfg.Port))
	if err != nil {
		slog.Error(fmt.Sprintf("failed to listen port: %s", cfg.Port), "err", err)
	}
	url := fmt.Sprintf("http://%s", l.Addr().String())
	slog.Info(fmt.Sprintf("start with %s", url))

	mux, err := router.New()
	if err != nil {
		slog.Error("failed to create router", "err", err)
		return err
	}

	s := server.New(l, mux)
	return s.Run(ctx)
}
