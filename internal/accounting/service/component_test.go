package service

import (
	"context"
	"github.com/google/uuid"
	"github.com/stretchr/testify/require"
	"log"
	"os"
	"testing"

	"github.com/noodlensk/task-tracker/internal/accounting/adapters"
	"github.com/noodlensk/task-tracker/internal/accounting/data/subscriber"
	"github.com/noodlensk/task-tracker/internal/accounting/ports/async"
	accountingAsyncPublisherClient "github.com/noodlensk/task-tracker/internal/common/clients/accounting/async/publisher"
	accountingCUDClient "github.com/noodlensk/task-tracker/internal/common/clients/accounting/cud/publisher"
	"github.com/noodlensk/task-tracker/internal/common/tests"
)

func TestEstimateTask(t *testing.T) {
	t.Parallel()

	ctx := context.Background()

	asyncPublisher := tests.NewAccountingAsyncPublisher(t)
	cudPublisher := tests.NewAccountingCUDPublisher(t)
	asyncSubscriber := tests.NewAccountingAsyncSubscriber(t)

	userToCreate := accountingCUDClient.UserCreated{
		Id:    uuid.New().String(),
		Name:  "Dmitry",
		Email: "some@email.com",
		Role:  "basic",
	}

	cudPublisher.CreateUser(t, userToCreate)

	taskToCreate := accountingCUDClient.TaskCreated{
		Id:         uuid.New().String(),
		Title:      "myTitle",
		AssignedTo: userToCreate.Id,
	}

	cudPublisher.CreateTask(t, taskToCreate)

	asyncPublisher.TaskCreated(t, accountingAsyncPublisherClient.TaskCreated{PublicId: taskToCreate.Id})

	estimatedTask, err := asyncSubscriber.WaitForTaskEstimated(ctx, taskToCreate.Id)
	require.NoError(t, err)

	require.True(t, float32(10) <= estimatedTask.AssignedPrice && float32(20) >= estimatedTask.AssignedPrice)
	require.True(t, float32(20) <= estimatedTask.CompetedPrice && float32(40) >= estimatedTask.CompetedPrice)
}

func TestChargeForAssignedTask(t *testing.T) {
	t.Parallel()

	ctx := context.Background()

	asyncPublisher := tests.NewAccountingAsyncPublisher(t)
	cudPublisher := tests.NewAccountingCUDPublisher(t)
	asyncSubscriber := tests.NewAccountingAsyncSubscriber(t)

	userToCreate := accountingCUDClient.UserCreated{
		Id:    uuid.New().String(),
		Name:  "Dmitry",
		Email: "some@email.com",
		Role:  "basic",
	}

	cudPublisher.CreateUser(t, userToCreate)

	taskToCreate := accountingCUDClient.TaskCreated{
		Id:          uuid.New().String(),
		Title:       "myTitle",
		Description: "my description",
		AssignedTo:  userToCreate.Id,
	}

	cudPublisher.CreateTask(t, taskToCreate)

	asyncPublisher.TaskAssigned(t, accountingAsyncPublisherClient.TaskAssigned{
		PublicId: taskToCreate.Id,
		UserId:   taskToCreate.AssignedTo,
	})

	userCharged, err := asyncSubscriber.WaitForUserCharged(ctx)
	require.NoError(t, err)

	require.Equal(t, userToCreate.Id, userCharged.UserUid)
}

func TestPayForFinishedTask(t *testing.T) {
	t.Parallel()

	ctx := context.Background()

	asyncPublisher := tests.NewAccountingAsyncPublisher(t)
	cudPublisher := tests.NewAccountingCUDPublisher(t)
	asyncSubscriber := tests.NewAccountingAsyncSubscriber(t)

	userToCreate := accountingCUDClient.UserCreated{
		Id:    uuid.New().String(),
		Name:  "Dmitry",
		Email: "some@email.com",
		Role:  "basic",
	}

	cudPublisher.CreateUser(t, userToCreate)

	taskToCreate := accountingCUDClient.TaskCreated{
		Id:          uuid.New().String(),
		Title:       "myTitle",
		Description: "",
		AssignedTo:  userToCreate.Id,
	}

	cudPublisher.CreateTask(t, taskToCreate)
	asyncPublisher.TaskCompleted(t, accountingAsyncPublisherClient.TaskCompleted{PublicId: taskToCreate.Id})

	userCharged, err := asyncSubscriber.WaitForUserPayed(ctx, userToCreate.Id)
	require.NoError(t, err)

	require.Equal(t, userToCreate.Id, userCharged.UserUid)
}

func startService() error {
	usersRepo := adapters.NewUserInMemoryRepository()
	taskRepo := adapters.NewTaskInMemoryRepository()

	app := NewComponentTestApplication(usersRepo, taskRepo)
	ctx := context.Background()

	asyncServer, err := tests.NewAsyncSubscriber()
	if err != nil {
		return err
	}

	if err := async.Register(async.NewAsyncServer(app), asyncServer.Router, asyncServer.Subscriber); err != nil {
		return err
	}

	if err := subscriber.Register(subscriber.NewAsyncServer(taskRepo, usersRepo), asyncServer.Router, asyncServer.Subscriber); err != nil {
		return err
	}
	go func() {
		err := asyncServer.Start(ctx) // TODO: wait for start
		if err != nil {
			panic(err)
		}
	}()

	asyncServer.Running()

	return nil
}

func TestMain(m *testing.M) {
	if err := startService(); err != nil {
		log.Printf("Failed to start service: %q\n", err.Error())
		os.Exit(1)
	}

	os.Exit(m.Run())
}
