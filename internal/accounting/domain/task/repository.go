package task

import "context"

type Repository interface {
	Get(ctx context.Context, uid string) (*Task, error)
	Update(ctx context.Context, uid string, updateFn func(ctx context.Context, task *Task) (*Task, error)) error
}
