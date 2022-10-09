package service

import (
	"context"
	http2 "github.com/noodlensk/task-tracker/internal/tasks/ports/http"
	"log"
	"net/http"
	"os"
	"testing"

	"github.com/go-chi/chi/v5"
	"github.com/google/uuid"
	"github.com/pkg/errors"
	"github.com/stretchr/testify/require"

	"github.com/noodlensk/task-tracker/internal/common/clients/tasks"
	"github.com/noodlensk/task-tracker/internal/common/server"
	"github.com/noodlensk/task-tracker/internal/common/tests"
)

func TestCreateTask(t *testing.T) {
	t.Parallel()

	token := tests.FakeAdminJWT(t, uuid.New().String())
	client := tests.NewTasksHTTPClient(t, token)

	taskToCreate := tasks.Task{
		Description: "some description",
		Title:       "My Task title",
	}
	client.CreateTask(t, taskToCreate)

	taskList := client.GetAllTasks(t)

	var taskCreated *tasks.Task

	for _, task := range taskList {
		if task.Title == taskToCreate.Title { //TODO: return UID of created task
			taskCreated = &task

			break
		}
	}

	require.NotNil(t, taskCreated)

	require.Equal(t, taskToCreate.Title, taskCreated.Title)
	require.Equal(t, taskToCreate.Description, taskCreated.Description)
	require.NotEmpty(t, taskCreated.Uid)
	// require.NotEmpty(t, taskCreated.AssignedTo) # TODO: fix it
	require.NotEmpty(t, taskCreated.CreatedAt)
	require.Equal(t, taskCreated.Status, tasks.NEW)
}

func startService() error {
	app := NewComponentTestApplication()
	ctx := context.Background()

	httpAddr := "127.0.0.1:8080"

	go server.RunHTTPServerOnAddr(ctx, httpAddr, tests.NewLogger(), func(router chi.Router) http.Handler {
		return http2.HandlerFromMux(http2.NewHTTPServer(app), router)
	})

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
