package cud

import (
	"context"

	"github.com/noodlensk/task-tracker/internal/accounting/domain/task"
)

type TaskUpdated struct {
	Task task.Task
}

type TaskUpdatedHandler struct {
	taskRepo task.Repository
}

func NewTaskUpdatedHandler(taskRepo task.Repository) TaskUpdatedHandler {
	return TaskUpdatedHandler{taskRepo: taskRepo}
}

func (h *TaskUpdatedHandler) Handle(ctx context.Context, cmd TaskUpdated) error {
	return h.taskRepo.Update(ctx, cmd.Task.UID(), func(ctx context.Context, t *task.Task) (*task.Task, error) {
		t.SetTitle(cmd.Task.Title())
		t.SetStatus(cmd.Task.Status())
		t.AssignToUser(cmd.Task.UID())

		return t, nil
	})
}
