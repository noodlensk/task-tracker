package adapters

import (
	"context"
	"errors"
	"fmt"
	"sync"

	"github.com/noodlensk/task-tracker/internal/tasks/domain/task"
	"github.com/noodlensk/task-tracker/internal/tasks/domain/user"
)

type TaskInMemoryRepository struct {
	tasks map[string]*task.Task

	lock sync.RWMutex
}

func NewTaskInMemoryRepository() *TaskInMemoryRepository {
	return &TaskInMemoryRepository{
		tasks: map[string]*task.Task{},
	}
}

func (r *TaskInMemoryRepository) AddTask(ctx context.Context, t *task.Task) error {
	r.lock.Lock()
	defer r.lock.Unlock()

	if _, ok := r.tasks[t.UID()]; ok {
		return fmt.Errorf("task with uid %q is already exist", t.UID())
	}

	r.tasks[t.UID()] = t

	return nil
}

func (r *TaskInMemoryRepository) GetTasksAssignedToUser(ctx context.Context, user user.User) ([]*task.Task, error) {
	r.lock.RLock()
	defer r.lock.RUnlock()

	var res []*task.Task

	for _, v := range r.tasks {
		if v.AssignedTo().UID != user.UID {
			continue
		}

		res = append(res, v)
	}

	return res, nil
}

func (r *TaskInMemoryRepository) GetAllTasks(ctx context.Context) ([]*task.Task, error) {
	r.lock.RLock()
	defer r.lock.RUnlock()

	var res []*task.Task

	for _, v := range r.tasks {
		res = append(res, v)
	}

	return res, nil
}

func (r *TaskInMemoryRepository) GetAllOpenTasks(ctx context.Context) ([]*task.Task, error) {
	r.lock.RLock()
	defer r.lock.RUnlock()

	var res []*task.Task

	for _, v := range r.tasks {
		if v.Status() == task.StatusDone {
			continue
		}

		res = append(res, v)
	}

	return res, nil
}

func (r *TaskInMemoryRepository) UpdateTask(ctx context.Context, taskUUID string, updateFn func(ctx context.Context, t *task.Task) (*task.Task, error)) error {
	r.lock.Lock()
	defer r.lock.Unlock()

	v, ok := r.tasks[taskUUID]
	if !ok {
		return errors.New("not found")
	}

	newVal, err := updateFn(ctx, v)
	if err != nil {
		return err
	}

	r.tasks[taskUUID] = newVal

	return nil
}
