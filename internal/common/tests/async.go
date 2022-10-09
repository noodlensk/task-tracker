package tests

import (
	"github.com/ThreeDotsLabs/watermill-kafka/v2/pkg/kafka"

	"github.com/noodlensk/task-tracker/internal/common/logs"
	"github.com/noodlensk/task-tracker/internal/common/server"
)

func NewAsyncSubscriber() (*server.WatermillServer, error) {
	sub, err := kafka.NewSubscriber(
		kafka.SubscriberConfig{
			Brokers:     []string{"localhost:9092"},
			Unmarshaler: kafka.DefaultMarshaler{},
		},
		logs.NewWatermillNopLogger(),
	)
	if err != nil {
		return nil, err
	}

	srv, err := server.NewWatermillServer(sub, logs.NewNopLogger())
	if err != nil {
		return nil, err
	}

	return srv, nil
}

func NewAsyncPublisher() (*kafka.Publisher, error) {
	return kafka.NewPublisher(
		kafka.PublisherConfig{
			Brokers:   []string{"localhost:9092"},
			Marshaler: kafka.DefaultMarshaler{},
		},
		logs.NewWatermillNopLogger(),
	)
}
