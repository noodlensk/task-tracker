package main

import (
	"context"
	"net/http"
	"os"
	"os/signal"

	"github.com/ThreeDotsLabs/watermill-kafka/v2/pkg/kafka"
	"github.com/go-chi/chi/v5"
	"go.uber.org/zap"
	"golang.org/x/sync/errgroup"

	"github.com/noodlensk/task-tracker/internal/common/logs"
	"github.com/noodlensk/task-tracker/internal/common/server"
	"github.com/noodlensk/task-tracker/internal/tasks/ports/async"
	httpTasks "github.com/noodlensk/task-tracker/internal/tasks/ports/http"
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

	app, err := service.NewApplication()
	if err != nil {
		return err
	}

	subscriber, err := kafka.NewSubscriber(
		kafka.SubscriberConfig{
			Brokers:     []string{"localhost:9092"},
			Unmarshaler: kafka.DefaultMarshaler{},
		},
		logs.NewWatermillLogger(logger.With("component", "watermill-subscriber-kafka")),
	)
	if err != nil {
		return err
	}

	watermillServer, err := server.NewWatermillServer(subscriber, logger.With("component", "watermill"))
	if err != nil {
		return err
	}

	if err := async.Register(async.NewAsyncServer(app), watermillServer.Router, watermillServer.Subscriber); err != nil {
		return err
	}

	wg, ctx := errgroup.WithContext(ctx)

	ctx, stop := signal.NotifyContext(ctx, os.Interrupt)
	defer stop()

	wg.Go(func() error {
		return server.RunHTTPServer(
			ctx,
			"localhost:8080",
			logger,
			func(router chi.Router) http.Handler {
				return httpTasks.HandlerFromMux(httpTasks.NewHTTPServer(app), router)
			},
		)
	})

	wg.Go(func() error { return watermillServer.Start(ctx) })

	return wg.Wait()
}
