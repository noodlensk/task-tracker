package command

import (
	"context"
	"fmt"
	"math/rand"

	"github.com/noodlensk/task-tracker/internal/tasks/domain/task"
	"github.com/noodlensk/task-tracker/internal/tasks/domain/user"
)

type CreateTask struct {
	Task task.Task
}

type CreateTaskHandler struct {
	tasksRepo task.Repository
	usersRepo user.Repository
}

func NewCreateTaskHandler(usersRepo user.Repository, tasksRepo task.Repository) CreateTaskHandler {
	return CreateTaskHandler{tasksRepo: tasksRepo, usersRepo: usersRepo}
}

func (h CreateTaskHandler) Handle(ctx context.Context, cmd CreateTask) error {
	basicUsers, err := h.usersRepo.GetUsers(ctx, "basic")
	if err != nil {
		return fmt.Errorf("get basic users: %w", err)
	}

	assignToUserNumber := rand.Intn(len(basicUsers)) //nolint:gosec

	cmd.Task.Assign(*basicUsers[assignToUserNumber])

	return h.tasksRepo.AddTask(ctx, &cmd.Task)
}
