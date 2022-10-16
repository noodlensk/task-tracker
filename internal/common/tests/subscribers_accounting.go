package tests

import (
	"context"
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/noodlensk/task-tracker/internal/common/clients/accounting/async/subscriber"
)

type AccountingAsyncSubscriber struct {
	taskEstimatedCh chan subscriber.TaskEstimated
	userChargedCh   chan subscriber.UserCharged
	userPayedCh     chan subscriber.UserPayed
}

func (a AccountingAsyncSubscriber) UserCharged(ctx context.Context, e subscriber.UserCharged) error {
	a.userChargedCh <- e

	return nil
}

func (a AccountingAsyncSubscriber) UserPayed(ctx context.Context, e subscriber.UserPayed) error {
	a.userPayedCh <- e

	return nil
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

func (a AccountingAsyncSubscriber) WaitForUserCharged(ctx context.Context) (*subscriber.UserCharged, error) {
	for {
		select {
		case <-ctx.Done():
			return nil, ctx.Err()
		case e := <-a.userChargedCh:
			return &e, nil
		}
	}
}

func (a AccountingAsyncSubscriber) WaitForUserPayed(ctx context.Context, userUID string) (*subscriber.UserPayed, error) {
	for {
		select {
		case <-ctx.Done():
			return nil, ctx.Err()
		case e := <-a.userPayedCh:
			if e.UserUid == userUID {
				return &e, nil
			}
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
		userPayedCh:     make(chan subscriber.UserPayed),
		userChargedCh:   make(chan subscriber.UserCharged),
	}

	require.NoError(t, subscriber.Register(client, srv.Router, srv.Subscriber))

	go func() {
		err := srv.Start(context.Background())
		require.NoError(t, err)
	}()

	<-srv.Router.Running()

	return client
}
