package cud

import (
	"context"

	"github.com/noodlensk/task-tracker/internal/accounting/domain/task"
)

type TaskCreated struct {
	Task task.Task
}

type TaskCreatedEventHandler struct {
	taskRepo task.Repository
}

func NewTaskCreatedEventHandler(taskRepo task.Repository) TaskCreatedEventHandler {
	return TaskCreatedEventHandler{taskRepo: taskRepo}
}

func (h *TaskCreatedEventHandler) Handle(ctx context.Context, cmd TaskCreated) error {
	return h.taskRepo.Create(ctx, &cmd.Task)
}
