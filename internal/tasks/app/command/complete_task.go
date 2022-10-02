package command

import (
	"context"
	"errors"
	"github.com/noodlensk/task-tracker/internal/tasks/domain/task"
	"github.com/noodlensk/task-tracker/internal/tasks/domain/user"
)

type CompleteTask struct {
	TaskUID string
	User    user.User
}

type CompleteTaskHandler struct {
	repo task.Repository
}

func NewCompleteTaskHandler(repo task.Repository) CompleteTaskHandler {
	return CompleteTaskHandler{repo: repo}
}

func (h CompleteTaskHandler) Handle(ctx context.Context, cmd CompleteTask) error {
	return h.repo.UpdateTask(ctx, cmd.TaskUID, func(ctx context.Context, t *task.Task) (*task.Task, error) {
		if t.AssignedTo().UID != cmd.User.UID {
			return nil, errors.New("only user assigned to the task could finish it")
		}

		t.Complete()

		return t, nil
	})
}
