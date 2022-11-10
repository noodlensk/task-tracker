package main

import (
	"context"
	"os"
	"os/signal"

	"github.com/ThreeDotsLabs/watermill-kafka/v2/pkg/kafka"
	"go.uber.org/zap"
	"golang.org/x/sync/errgroup"

	"github.com/noodlensk/task-tracker/internal/analytics/adapters"
	"github.com/noodlensk/task-tracker/internal/analytics/data/subscriber"
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
	transactionsRepo := adapters.NewTransactionInMemoryRepository()

	watermillServer, err := server.NewWatermillServer(sub, logger.With("component", "watermill"))
	if err != nil {
		return err
	}

	if err := subscriber.Register(subscriber.NewAsyncServer(transactionsRepo, usersRepo), watermillServer.Router, watermillServer.Subscriber); err != nil {
		return err
	}

	wg, ctx := errgroup.WithContext(ctx)

	ctx, stop := signal.NotifyContext(ctx, os.Interrupt)
	defer stop()

	wg.Go(func() error { return watermillServer.Start(ctx) })

	return wg.Wait()
}
