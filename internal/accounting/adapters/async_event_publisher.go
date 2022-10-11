package adapters

import (
	"context"
	"fmt"
	"time"

	"github.com/ThreeDotsLabs/watermill/message"

	"github.com/noodlensk/task-tracker/internal/accounting/app/command"
)

type AsyncEventPublisher struct {
	client *PublisherClient
}

func (a AsyncEventPublisher) PayedForFinishedTask(ctx context.Context, e command.PayedForFinishedTask) {
	// TODO: handle
}

func (a AsyncEventPublisher) TaskPriceEstimated(ctx context.Context, e command.TaskPriceEstimated) {
	err := a.client.TaskEstimated(ctx, TaskEstimated{
		Id:            e.TaskUID,
		AssignedPrice: e.PriceAssigned,
		CompetedPrice: e.PriceFinished,
		SentAt:        time.Now(),
	})

	if err != nil {
		fmt.Println(err) // TODO: fix this
	}
}

func NewAsyncEventPublisher(pub message.Publisher) *AsyncEventPublisher {
	return &AsyncEventPublisher{
		client: NewPublisherClient(pub),
	}
}
