package query

import (
	"context"

	"github.com/noodlensk/task-tracker/internal/tasks/domain/task"
)

type AllTasks struct {
	Limit  int
	Offset int
}

type AllTasksHandler struct {
	repo task.Repository
}

func NewAllTasksHandler(repo task.Repository) AllTasksHandler {
	return AllTasksHandler{repo: repo}
}

func (h AllTasksHandler) Handle(ctx context.Context, tasks AllTasks) ([]*task.Task, error) {
	return h.repo.GetAllTasks(ctx)
}
