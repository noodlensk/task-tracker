package main

import (
	"context"
	"net/http"
	"os"
	"os/signal"

	"github.com/ThreeDotsLabs/watermill-kafka/v2/pkg/kafka"
	"github.com/go-chi/chi/v5"
	"go.uber.org/zap"

	"github.com/noodlensk/task-tracker/internal/common/logs"
	"github.com/noodlensk/task-tracker/internal/common/server"
	"github.com/noodlensk/task-tracker/internal/users/ports"
	"github.com/noodlensk/task-tracker/internal/users/service"
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

	publisher, err := kafka.NewPublisher(
		kafka.PublisherConfig{
			Brokers:   []string{"localhost:9092"},
			Marshaler: kafka.DefaultMarshaler{},
		},
		logs.NewWatermillLogger(logger.With("component", "watermill-publisher-kafka")),
	)
	if err != nil {
		return err
	}

	app, err := service.NewApplication(publisher)
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
