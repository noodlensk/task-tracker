package adapters

import (
	"context"
	"errors"
	"sync"

	"github.com/noodlensk/task-tracker/internal/accounting/domain/task"
)

type TaskInMemoryRepository struct {
	tasks map[string]*task.Task

	lock sync.RWMutex
}

func (r *TaskInMemoryRepository) Create(ctx context.Context, task *task.Task) error {
	r.lock.Lock()
	defer r.lock.Unlock()

	r.tasks[task.UID()] = task

	return nil
}

func (r *TaskInMemoryRepository) Get(ctx context.Context, uid string) (*task.Task, error) {
	r.lock.RLock()
	defer r.lock.RUnlock()

	v, ok := r.tasks[uid]
	if !ok {
		return nil, errors.New("not found")
	}

	return v, nil
}

func (r *TaskInMemoryRepository) Update(ctx context.Context, uid string, updateFn func(ctx context.Context, task *task.Task) (*task.Task, error)) error {
	r.lock.Lock()
	defer r.lock.Unlock()

	v, ok := r.tasks[uid]
	if !ok {
		return errors.New("not found")
	}

	newVal, err := updateFn(ctx, v)
	if err != nil {
		return err
	}

	r.tasks[uid] = newVal

	return nil
}

func NewTaskInMemoryRepository() *TaskInMemoryRepository {
	return &TaskInMemoryRepository{
		tasks: map[string]*task.Task{},
	}
}
