package cud

import (
	"context"

	"github.com/noodlensk/task-tracker/internal/accounting/domain/task"
)

type TaskCreated struct {
	Task task.Task
}

type TaskCreatedHandler struct {
	taskRepo task.Repository
}

func NewTaskCreatedHandler(taskRepo task.Repository) TaskCreatedHandler {
	return TaskCreatedHandler{taskRepo: taskRepo}
}

func (h *TaskCreatedHandler) Handle(ctx context.Context, cmd TaskCreated) error {
	return h.taskRepo.Create(ctx, &cmd.Task)
}
