package main

import (
	"context"
	"net/http"
	"os"
	"os/signal"

	"github.com/go-chi/chi/v5"
	"go.uber.org/zap"

	"github.com/noodlensk/task-tracker/internal/common/logs"
	"github.com/noodlensk/task-tracker/internal/common/server"
	"github.com/noodlensk/task-tracker/internal/tasks/ports"
	"github.com/noodlensk/task-tracker/internal/tasks/service"
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

	articleApp, err := service.NewApplication()
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
