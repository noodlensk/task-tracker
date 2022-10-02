package main

import (
	"context"
	"github.com/go-chi/chi/v5"
	"github.com/noodlensk/task-tracker/internal/common/logs"
	"github.com/noodlensk/task-tracker/internal/common/server"
	"github.com/noodlensk/task-tracker/internal/tasks/ports"
	"github.com/noodlensk/task-tracker/internal/tasks/service"
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

	articleApp, err := service.NewApplication(logger.With("app", "task"))
	if err != nil {
		return err
	}

	ctx, stop := signal.NotifyContext(ctx, os.Interrupt)
	defer stop()

	return server.RunHTTPServer(
		ctx,
		"localhost:8080",
		logger,
		func(router chi.Router) http.Handler {
			return ports.HandlerFromMux(ports.NewHTTPServer(articleApp), router)
		},
	)
}