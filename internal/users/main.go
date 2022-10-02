package main

import (
	"context"
	"github.com/go-chi/chi/v5"
	"github.com/noodlensk/task-tracker/internal/common/logs"
	"github.com/noodlensk/task-tracker/internal/common/server"
	"github.com/noodlensk/task-tracker/internal/users/ports"
	"github.com/noodlensk/task-tracker/internal/users/service"

	"go.uber.org/zap"
	"net/http"
	"os"
	"os/signal"
)

func main() {
	logger := logs.NewLogger()

	if err := run(logger); err != nil {
		if err != context.Canceled {
			logger.Fatal(err)
		}
	}
}

func run(logger *zap.SugaredLogger) error {
	ctx := context.Background()

	app, err := service.NewApplication(logger.With("app", "user"))
	if err != nil {
		return err
	}

	ctx, stop := signal.NotifyContext(ctx, os.Interrupt)
	defer stop()

	return server.RunHTTPServer(
		ctx,
		"localhost:8081",
		logger,
		func(router chi.Router) http.Handler {
			return ports.HandlerFromMux(ports.NewHTTPServer(app), router)
		},
	)
}
