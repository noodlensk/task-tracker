package command

import (
	"context"
	"fmt"
	"math/rand"

	"github.com/noodlensk/task-tracker/internal/tasks/domain/task"
	"github.com/noodlensk/task-tracker/internal/tasks/domain/user"
)

type ReAssignAllTasks struct {
	User user.User
}

type ReAssignAllTasksHandler struct {
	userRepo user.Repository
	taskRepo task.Repository
}

func NewReAssignAllTasksHandler(userRepo user.Repository, taskRepo task.Repository) ReAssignAllTasksHandler {
	return ReAssignAllTasksHandler{userRepo: userRepo, taskRepo: taskRepo}
}

func (h ReAssignAllTasksHandler) Handle(ctx context.Context, cmd ReAssignAllTasks) error {
	basicUsers, err := h.userRepo.GetUsers(ctx, "basic")
	if err != nil {
		return fmt.Errorf("get basic users: %w", err)
	}

	allOpenTasks, err := h.taskRepo.GetAllOpenTasks(ctx)
	if err != nil {
		return fmt.Errorf("get open tasks: %w", err)
	}

	for _, t := range allOpenTasks {
		assignToUserNumber := rand.Intn(len(basicUsers)) //nolint:gosec

		err := h.taskRepo.UpdateTask(ctx, t.UID(), func(ctx context.Context, t *task.Task) (*task.Task, error) {
			t.Assign(*basicUsers[assignToUserNumber])

			return t, nil
		})
		if err != nil {
			return fmt.Errorf("assign task %q to user %q: %w", t.UID(), basicUsers[assignToUserNumber].UID, err)
		}
	}

	return nil
}
