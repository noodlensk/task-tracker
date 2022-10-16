package cud

import (
	"context"

	"github.com/noodlensk/task-tracker/internal/accounting/domain/task"
)

type TaskUpdated struct {
	Task task.Task
}

type TaskUpdatedEventHandler struct {
	taskRepo task.Repository
}

func NewTaskUpdatedEventHandler(taskRepo task.Repository) TaskUpdatedEventHandler {
	return TaskUpdatedEventHandler{taskRepo: taskRepo}
}

func (h *TaskUpdatedEventHandler) Handle(ctx context.Context, cmd TaskUpdated) error {
	return h.taskRepo.Update(ctx, cmd.Task.UID(), func(ctx context.Context, t *task.Task) (*task.Task, error) {
		t.SetTitle(cmd.Task.Title())
		t.SetStatus(cmd.Task.Status())
		t.AssignToUser(cmd.Task.AssignedToUserUID())

		return t, nil
	})
}
