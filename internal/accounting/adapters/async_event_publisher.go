package adapters

import (
	"context"

	"github.com/ThreeDotsLabs/watermill/message"
	"go.uber.org/zap"

	"github.com/noodlensk/task-tracker/internal/accounting/app/command"
)

type AsyncEventPublisher struct {
	client *PublisherClient
	logger *zap.SugaredLogger
}

func (a AsyncEventPublisher) UserChargedForAssignedTask(ctx context.Context, e command.ChargedForAssignedTask) {
	err := a.client.UserCharged(ctx, UserCharged{
		UserUid: e.UserUID,
		Reason:  "", // TODO: fix this
		Amount:  e.Amount,
	})
	if err != nil {
		a.logger.Errorw("Failed to publish event", "err", err)
	}
}

func (a AsyncEventPublisher) PayedForFinishedTask(ctx context.Context, e command.PayedForFinishedTask) {
	err := a.client.UserPayed(ctx, UserPayed{
		UserUid: e.UserUID,
		Reason:  "", // TODO: fix this
		Amount:  e.Amount,
	})
	if err != nil {
		a.logger.Errorw("Failed to publish event", "err", err)
	}
}

func (a AsyncEventPublisher) TaskPriceEstimated(ctx context.Context, e command.TaskPriceEstimated) {
	err := a.client.TaskEstimated(ctx, TaskEstimated{
		Id:            e.TaskUID,
		AssignedPrice: e.PriceAssigned,
		CompetedPrice: e.PriceFinished,
	})
	if err != nil {
		a.logger.Errorw("Failed to publish event", "err", err)
	}
}

func NewAsyncEventPublisher(pub message.Publisher, logger *zap.SugaredLogger) *AsyncEventPublisher {
	return &AsyncEventPublisher{
		client: NewPublisherClient(pub),
		logger: logger,
	}
}
