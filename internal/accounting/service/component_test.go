package service

import (
	"context"
	"github.com/noodlensk/task-tracker/internal/accounting/ports/async"
	accountingAsyncPublisherClient "github.com/noodlensk/task-tracker/internal/common/clients/accounting/async/publisher"
	"github.com/noodlensk/task-tracker/internal/common/tests"
	"github.com/stretchr/testify/require"
	"log"
	"os"
	"testing"
)

func TestEstimateTask(t *testing.T) {
	t.Parallel()

	ctx := context.Background()

	asyncPublisher := tests.NewAccountingAsyncPublisher(t)

	asyncSubscriber := tests.NewAccountingAsyncSubscriber(t)

	userToCreate := accountingAsyncPublisherClient.UserCreated{
		Id:    "myUID",
		Name:  "Dmitry",
		Email: "some@email.com",
		Role:  "basic",
	}

	asyncPublisher.UserCreated(t, userToCreate)

	taskToCreate := accountingAsyncPublisherClient.TaskCreated{
		Id:          "myUID",
		Title:       "myTitle",
		Description: "",
		AssignedTo:  userToCreate.Id,
	}

	asyncPublisher.TaskCreated(t, taskToCreate)

	estimatedTask, err := asyncSubscriber.WaitForTaskEstimated(ctx)
	require.NoError(t, err)

	require.True(t, float32(10) <= estimatedTask.AssignedPrice && float32(20) >= estimatedTask.AssignedPrice)
	require.True(t, float32(20) <= estimatedTask.CompetedPrice && float32(40) >= estimatedTask.CompetedPrice)
}

func startService() error {
	app := NewComponentTestApplication()
	ctx := context.Background()

	asyncServer, err := tests.NewAsyncSubscriber()
	if err != nil {
		return err
	}

	if err := async.Register(async.NewAsyncServer(app), asyncServer.Router, asyncServer.Subscriber); err != nil {
		return err
	}

	go asyncServer.Start(ctx) // TODO: wait for start

	return nil
}

func TestMain(m *testing.M) {
	if err := startService(); err != nil {
		log.Printf("Failed to start service: %q\n", err.Error())
		os.Exit(1)
	}

	os.Exit(m.Run())
}
