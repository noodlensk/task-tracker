package service

import (
	"context"
	"log"
	"os"
	"testing"

	"github.com/google/uuid"
	"github.com/stretchr/testify/require"

	"github.com/noodlensk/task-tracker/internal/accounting/ports/async"
	accountingAsyncPublisherClient "github.com/noodlensk/task-tracker/internal/common/clients/accounting/async/publisher"
	"github.com/noodlensk/task-tracker/internal/common/tests"
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

func TestChargeForAssignedTask(t *testing.T) {
	t.Parallel()

	ctx := context.Background()

	asyncPublisher := tests.NewAccountingAsyncPublisher(t)
	asyncSubscriber := tests.NewAccountingAsyncSubscriber(t)

	userToCreate := accountingAsyncPublisherClient.UserCreated{
		Id:    uuid.New().String(),
		Name:  "Dmitry",
		Email: "some@email.com",
		Role:  "basic",
	}

	asyncPublisher.UserCreated(t, userToCreate)

	taskToCreate := accountingAsyncPublisherClient.TaskCreated{
		Id:          uuid.New().String(),
		Title:       "myTitle",
		Description: "my description",
		AssignedTo:  userToCreate.Id,
	}

	asyncPublisher.TaskCreated(t, taskToCreate)
	asyncPublisher.TaskAssigned(t, accountingAsyncPublisherClient.TaskAssigned{
		Id:          taskToCreate.Id,
		Title:       taskToCreate.Title,
		Description: taskToCreate.Description,
		AssignedTo:  taskToCreate.AssignedTo,
	})

	userCharged, err := asyncSubscriber.WaitForUserCharged(ctx)
	require.NoError(t, err)

	require.Equal(t, userToCreate.Id, userCharged.UserUid)
}

func TestPayForFinishedTask(t *testing.T) {
	t.Parallel()

	ctx := context.Background()

	asyncPublisher := tests.NewAccountingAsyncPublisher(t)
	asyncSubscriber := tests.NewAccountingAsyncSubscriber(t)

	userToCreate := accountingAsyncPublisherClient.UserCreated{
		Id:    uuid.New().String(),
		Name:  "Dmitry",
		Email: "some@email.com",
		Role:  "basic",
	}

	asyncPublisher.UserCreated(t, userToCreate)

	taskToCreate := accountingAsyncPublisherClient.TaskCreated{
		Id:          uuid.New().String(),
		Title:       "myTitle",
		Description: "",
		AssignedTo:  userToCreate.Id,
	}

	asyncPublisher.TaskCreated(t, taskToCreate)
	asyncPublisher.TaskCompleted(t, accountingAsyncPublisherClient.TaskCompleted{
		Id:          taskToCreate.Id,
		Title:       taskToCreate.Title,
		Description: taskToCreate.Description,
		AssignedTo:  taskToCreate.AssignedTo,
	})

	userCharged, err := asyncSubscriber.WaitForUserPayed(ctx, userToCreate.Id)
	require.NoError(t, err)

	require.Equal(t, userToCreate.Id, userCharged.UserUid)
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

	go func() {
		err := asyncServer.Start(ctx) // TODO: wait for start
		if err != nil {
			panic(err)
		}
	}()

	return nil
}

func TestMain(m *testing.M) {
	if err := startService(); err != nil {
		log.Printf("Failed to start service: %q\n", err.Error())
		os.Exit(1)
	}

	os.Exit(m.Run())
}
