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
	"github.com/noodlensk/task-tracker/internal/tasks/adapters"
	"github.com/noodlensk/task-tracker/internal/tasks/data/publisher"
	"github.com/noodlensk/task-tracker/internal/tasks/data/subscriber"
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

	sub, err := kafka.NewSubscriber(
		kafka.SubscriberConfig{
			Brokers:     []string{"localhost:9092"},
			Unmarshaler: kafka.DefaultMarshaler{},
		},
		logs.NewWatermillLogger(logger.With("component", "watermill-subscriber-kafka")),
	)
	if err != nil {
		return err
	}

	watermillServer, err := server.NewWatermillServer(sub, logger.With("component", "watermill"))
	if err != nil {
		return err
	}

	pub, err := kafka.NewPublisher(
		kafka.PublisherConfig{
			Brokers:   []string{"localhost:9092"},
			Marshaler: kafka.DefaultMarshaler{},
		},
		logs.NewWatermillLogger(logger.With("component", "watermill-publisher-kafka")),
	)
	if err != nil {
		return err
	}

	taskRepo := adapters.NewTaskInMemoryRepository(publisher.NewPublisherClient(pub))
	userRepo := adapters.NewUserInMemoryRepository()

	if err := subscriber.Register(subscriber.NewAsyncServer(userRepo), watermillServer.Router, watermillServer.Subscriber); err != nil {
		return err
	}

	app, err := service.NewApplication(userRepo, taskRepo)
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
