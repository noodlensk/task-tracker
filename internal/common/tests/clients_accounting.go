package tests

import (
	"context"
	"testing"

	"github.com/stretchr/testify/require"

	accountingAsyncPublisherClient "github.com/noodlensk/task-tracker/internal/common/clients/accounting/async/publisher"
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

func (c AccountingAsyncClient) UserCreated(t *testing.T, u accountingAsyncPublisherClient.UserCreated) {
	t.Helper()

	ctx := context.Background()

	err := c.client.UserCreated(ctx, u)

	require.NoError(t, err)
}

func (c AccountingAsyncClient) TaskCreated(t *testing.T, task accountingAsyncPublisherClient.TaskCreated) {
	t.Helper()

	ctx := context.Background()

	err := c.client.TaskCreated(ctx, task)

	require.NoError(t, err)
}
