package tests

import (
	"context"
	"testing"

	"github.com/stretchr/testify/require"

	accountingAsyncPublisherClient "github.com/noodlensk/task-tracker/internal/common/clients/accounting/async/publisher"
	"github.com/noodlensk/task-tracker/internal/common/clients/accounting/cud/publisher"
)

type AccountingAsyncClient struct {
	client *accountingAsyncPublisherClient.PublisherClient
}

func NewAccountingAsyncPublisher(t *testing.T) AccountingAsyncClient {
	t.Helper()

	pub, err := NewAsyncPublisher()
	require.NoError(t, err)

	return AccountingAsyncClient{
		client: accountingAsyncPublisherClient.NewPublisherClient(pub),
	}
}

func (c AccountingAsyncClient) TaskAssigned(t *testing.T, task accountingAsyncPublisherClient.TaskAssigned) {
	t.Helper()

	ctx := context.Background()

	err := c.client.TaskAssigned(ctx, task)

	require.NoError(t, err)
}

func (c AccountingAsyncClient) TaskCreated(t *testing.T, task accountingAsyncPublisherClient.TaskCreated) {
	t.Helper()

	ctx := context.Background()

	err := c.client.TaskCreated(ctx, task)

	require.NoError(t, err)
}

func (c AccountingAsyncClient) TaskCompleted(t *testing.T, task accountingAsyncPublisherClient.TaskCompleted) {
	t.Helper()

	ctx := context.Background()

	err := c.client.TaskCompleted(ctx, task)

	require.NoError(t, err)
}

type AccountingCUDPublisher struct {
	client *publisher.PublisherClient
}

func NewAccountingCUDPublisher(t *testing.T) AccountingCUDPublisher {
	t.Helper()

	pub, err := NewAsyncPublisher()
	require.NoError(t, err)

	return AccountingCUDPublisher{
		client: publisher.NewPublisherClient(pub),
	}
}

func (c AccountingCUDPublisher) CreateUser(t *testing.T, u publisher.UserCreated) {
	t.Helper()

	ctx := context.Background()

	err := c.client.UserCreated(ctx, u)

	require.NoError(t, err)
}

func (c AccountingCUDPublisher) CreateTask(t *testing.T, tsk publisher.TaskCreated) {
	t.Helper()

	ctx := context.Background()

	err := c.client.TaskCreated(ctx, tsk)

	require.NoError(t, err)
}
