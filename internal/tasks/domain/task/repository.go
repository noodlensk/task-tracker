package task

import (
	"context"

	"github.com/noodlensk/task-tracker/internal/tasks/domain/user"
)

type Repository interface {
	AddTask(ctx context.Context, t *Task) error
	GetTasksAssignedToUser(ctx context.Context, user user.User) ([]*Task, error)
	GetAllTasks(ctx context.Context) ([]*Task, error)
	GetAllOpenTasks(ctx context.Context) ([]*Task, error)

	UpdateTask(
		ctx context.Context,
		taskUUID string,
		updateFn func(ctx context.Context, t *Task) (*Task, error),
	) error
}
