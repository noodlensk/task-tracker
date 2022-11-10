package service

import (
	"context"
	"log"
	"net/http"
	"os"
	"testing"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/google/uuid"
	"github.com/pkg/errors"
	"github.com/stretchr/testify/require"

	tasksCUDClient "github.com/noodlensk/task-tracker/internal/common/clients/tasks/cud/publisher"
	tasksHTTPClient "github.com/noodlensk/task-tracker/internal/common/clients/tasks/http"
	"github.com/noodlensk/task-tracker/internal/common/server"
	"github.com/noodlensk/task-tracker/internal/common/tests"
	"github.com/noodlensk/task-tracker/internal/tasks/adapters"
	"github.com/noodlensk/task-tracker/internal/tasks/data/publisher"
	"github.com/noodlensk/task-tracker/internal/tasks/data/subscriber"
	tasksAsyncServer "github.com/noodlensk/task-tracker/internal/tasks/ports/async"
	tasksHTTPServer "github.com/noodlensk/task-tracker/internal/tasks/ports/http"
)

func TestCreateTask(t *testing.T) {
	t.Parallel()

	token := tests.FakeAdminJWT(t, uuid.New().String())
	httpClient := tests.NewTasksHTTPClient(t, token)

	asyncClient := tests.NewTasksCUDClient(t)

	userToCreate := tasksCUDClient.UserCreated{
		Id:    "myUID",
		Name:  "Dmitry",
		Email: "some@email.com",
		Role:  "basic",
	}

	asyncClient.CreateUser(t, userToCreate)

	time.Sleep(time.Second * 1) // TODO: replace it with more stable solution

	taskToCreate := tasksHTTPClient.Task{
		Description: "some description",
		Title:       "My Task title",
	}
	httpClient.CreateTask(t, taskToCreate)

	taskList := httpClient.GetAllTasks(t)

	var taskCreated *tasksHTTPClient.Task

	for _, task := range taskList {
		if task.Title == taskToCreate.Title { // TODO: return UID of created task
			taskCreated = &task

			break
		}
	}

	require.NotNil(t, taskCreated)

	require.Equal(t, taskToCreate.Title, taskCreated.Title)
	require.Equal(t, taskToCreate.Description, taskCreated.Description)
	require.NotEmpty(t, taskCreated.Uid)
	require.Equal(t, userToCreate.Id, *taskCreated.AssignedTo)
	require.NotEmpty(t, taskCreated.CreatedAt)
	require.Equal(t, tasksHTTPClient.NEW, *taskCreated.Status)
}

func startService() error {
	asyncPub, err := tests.NewAsyncPublisher()
	if err != nil {
		panic(err)
	}

	pub := publisher.NewPublisherClient(asyncPub)

	taskRepo := adapters.NewTaskInMemoryRepository(pub)

	userRepo := adapters.NewUserInMemoryRepository()

	app := NewComponentTestApplication(userRepo, taskRepo)
	ctx := context.Background()

	httpAddr := "127.0.0.1:8080"

	asyncServer, err := tests.NewAsyncSubscriber()
	if err != nil {
		return err
	}

	if err := tasksAsyncServer.Register(tasksAsyncServer.NewAsyncServer(app), asyncServer.Router, asyncServer.Subscriber); err != nil {
		return err
	}

	if err := subscriber.Register(subscriber.NewAsyncServer(userRepo), asyncServer.Router, asyncServer.Subscriber); err != nil {
		return err
	}

	go server.RunHTTPServerOnAddr(ctx, httpAddr, tests.NewLogger(), func(router chi.Router) http.Handler {
		return tasksHTTPServer.HandlerFromMux(tasksHTTPServer.NewHTTPServer(app), router)
	})

	go asyncServer.Start(ctx)

	asyncServer.Running()

	ok := tests.WaitForPort(httpAddr)
	if !ok {
		return errors.Errorf("Timed out waiting for tasks HTTP to come up")
	}

	return nil
}

func TestMain(m *testing.M) {
	if err := startService(); err != nil {
		log.Printf("Failed to start service: %q\n", err.Error())
		os.Exit(1)
	}

	os.Exit(m.Run())
}
