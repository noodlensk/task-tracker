package query

import (
	"context"

	"github.com/noodlensk/task-tracker/internal/tasks/domain/task"
	"github.com/noodlensk/task-tracker/internal/tasks/domain/user"
)

type TasksForUser struct {
	User   user.User
	Limit  int
	Offset int
}

type TasksForUserHandler struct {
	repo task.Repository
}

func NewTasksForUserHandler(repo task.Repository) TasksForUserHandler {
	return TasksForUserHandler{repo: repo}
}

func (h TasksForUserHandler) Handle(ctx context.Context, q TasksForUser) ([]*task.Task, error) {
	return h.repo.GetTasksAssignedToUser(ctx, q.User)
}
