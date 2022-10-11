package tests

import (
	"context"
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/noodlensk/task-tracker/internal/common/clients/accounting/async/subscriber"
)

type AccountingAsyncSubscriber struct {
	taskEstimatedCh chan subscriber.TaskEstimated
}

func (a AccountingAsyncSubscriber) TaskEstimated(ctx context.Context, e subscriber.TaskEstimated) error {
	a.taskEstimatedCh <- e

	return nil
}

func (a AccountingAsyncSubscriber) WaitForTaskEstimated(ctx context.Context) (*subscriber.TaskEstimated, error) {
	for {
		select {
		case <-ctx.Done():
			return nil, ctx.Err()
		case e := <-a.taskEstimatedCh:
			return &e, nil
		}
	}
}

func (a AccountingAsyncSubscriber) UserCreated(ctx context.Context, e subscriber.UserCreated) error {
	return nil
}

func (a AccountingAsyncSubscriber) UserUpdated(ctx context.Context, e subscriber.UserUpdated) error {
	return nil
}

func (a AccountingAsyncSubscriber) TaskAssigned(ctx context.Context, e subscriber.TaskAssigned) error {
	return nil
}

func (a AccountingAsyncSubscriber) TaskCreated(ctx context.Context, e subscriber.TaskCreated) error {
	return nil
}

func (a AccountingAsyncSubscriber) TaskCompleted(ctx context.Context, e subscriber.TaskCompleted) error {
	return nil
}

func NewAccountingAsyncSubscriber(t *testing.T) *AccountingAsyncSubscriber {
	t.Helper()

	srv, err := NewAsyncSubscriber()
	require.NoError(t, err)

	client := &AccountingAsyncSubscriber{
		taskEstimatedCh: make(chan subscriber.TaskEstimated),
	}

	require.NoError(t, subscriber.Register(client, srv.Router, srv.Subscriber))

	go func() {
		err := srv.Start(context.Background())
		require.NoError(t, err)
	}()

	<-srv.Router.Running()

	return client
}
