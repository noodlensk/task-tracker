package command

import (
	"context"

	"github.com/noodlensk/task-tracker/internal/tasks/domain/task"
)

type CreateTask struct {
	Task task.Task
}

type CreateTaskHandler struct {
	repo task.Repository
}

func NewCreateTaskHandler(repo task.Repository) CreateTaskHandler {
	return CreateTaskHandler{repo: repo}
}

func (h CreateTaskHandler) Handle(ctx context.Context, cmd CreateTask) error {
	return h.repo.AddTask(ctx, &cmd.Task)
}
