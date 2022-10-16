package main

import (
	"context"
	"os"
	"os/signal"

	"github.com/ThreeDotsLabs/watermill-kafka/v2/pkg/kafka"
	"go.uber.org/zap"
	"golang.org/x/sync/errgroup"

	"github.com/noodlensk/task-tracker/internal/accounting/adapters"
	"github.com/noodlensk/task-tracker/internal/accounting/data/subscriber"
	"github.com/noodlensk/task-tracker/internal/accounting/ports/async"
	"github.com/noodlensk/task-tracker/internal/accounting/service"
	"github.com/noodlensk/task-tracker/internal/common/logs"
	"github.com/noodlensk/task-tracker/internal/common/server"
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

	usersRepo := adapters.NewUserInMemoryRepository()
	taskRepo := adapters.NewTaskInMemoryRepository()

	app, err := service.NewApplication(usersRepo, taskRepo, pub, logger)
	if err != nil {
		return err
	}

	watermillServer, err := server.NewWatermillServer(sub, logger.With("component", "watermill"))
	if err != nil {
		return err
	}

	if err := async.Register(async.NewAsyncServer(app), watermillServer.Router, watermillServer.Subscriber); err != nil {
		return err
	}

	if err := subscriber.Register(subscriber.NewAsyncServer(taskRepo, usersRepo), watermillServer.Router, watermillServer.Subscriber); err != nil {
		return err
	}

	wg, ctx := errgroup.WithContext(ctx)

	ctx, stop := signal.NotifyContext(ctx, os.Interrupt)
	defer stop()

	wg.Go(func() error { return watermillServer.Start(ctx) })

	return wg.Wait()
}
