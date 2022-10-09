package server

import (
	"context"

	"github.com/ThreeDotsLabs/watermill/message"
	"go.uber.org/zap"

	"github.com/noodlensk/task-tracker/internal/common/logs"
)

type WatermillServer struct {
	Router     *message.Router
	Subscriber message.Subscriber

	logger *zap.SugaredLogger
}

func NewWatermillServer(sub message.Subscriber, logger *zap.SugaredLogger) (*WatermillServer, error) {
	router, err := message.NewRouter(message.RouterConfig{}, logs.NewWatermillLogger(logger))
	if err != nil {
		return nil, err
	}

	return &WatermillServer{
		Router:     router,
		Subscriber: sub,
		logger:     logger,
	}, nil
}

func (s WatermillServer) Start(ctx context.Context) error {
	s.logger.Info("Starting watermill server")

	return s.Router.Run(ctx)
}
