package tests

import (
	"context"
	"fmt"
	"net/http"
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/noodlensk/task-tracker/internal/common/clients/tasks"
)

type TasksHTTPClient struct {
	client *tasks.ClientWithResponses
}

func NewTasksHTTPClient(t *testing.T, token string) TasksHTTPClient {
	t.Helper()

	addr := "localhost:8080"

	t.Log("Trying tasks http:", addr)
	ok := WaitForPort(addr)
	require.True(t, ok, "Tasks HTTP timed out")

	url := fmt.Sprintf("http://%v/api", addr)

	client, err := tasks.NewClientWithResponses(
		url,
		tasks.WithRequestEditorFn(authorizationBearer(token)),
	)
	require.NoError(t, err)

	return TasksHTTPClient{client: client}
}

func (c TasksHTTPClient) CreateTask(t *testing.T, task tasks.Task) {
	t.Helper()

	resp, err := c.client.CreateTask(context.Background(), task)

	require.NoError(t, err)
	require.NoError(t, resp.Body.Close())

	require.Equal(t, http.StatusCreated, resp.StatusCode)
}

func (c TasksHTTPClient) GetAllTasks(t *testing.T) []tasks.Task {
	t.Helper()

	resp, err := c.client.GetTasksWithResponse(context.Background(), &tasks.GetTasksParams{Limit: 100, Offset: 0})

	require.NoError(t, err)
	require.Equal(t, http.StatusOK, resp.StatusCode())

	rp := *resp.JSON200

	return rp.Tasks
}
