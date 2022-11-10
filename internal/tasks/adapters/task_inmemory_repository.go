package adapters

import (
	"context"
	"errors"
	"fmt"
	"sync"

	"github.com/noodlensk/task-tracker/internal/tasks/data/publisher"
	"github.com/noodlensk/task-tracker/internal/tasks/domain/task"
	"github.com/noodlensk/task-tracker/internal/tasks/domain/user"
)

type TaskInMemoryRepository struct {
	tasks map[string]*task.Task

	publisher *publisher.PublisherClient

	lock sync.RWMutex
}

func NewTaskInMemoryRepository(pub *publisher.PublisherClient) *TaskInMemoryRepository {
	return &TaskInMemoryRepository{
		tasks:     map[string]*task.Task{},
		publisher: pub,
	}
}

func (r *TaskInMemoryRepository) AddTask(ctx context.Context, t *task.Task) error {
	r.lock.Lock()
	defer r.lock.Unlock()

	if _, ok := r.tasks[t.UID()]; ok {
		return fmt.Errorf("task with uid %q is already exist", t.UID())
	}

	r.tasks[t.UID()] = t

	if err := r.publisher.TaskCreated(ctx, publisher.TaskCreated{
		Id:          t.UID(),
		Title:       t.Title(),
		Description: t.Description(),
		AssignedTo:  t.AssignedTo().UID,
		CreatedBy:   t.CreatedBy().UID,
		Status:      string(t.Status()),
	}); err != nil {
		return err
	}

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

	if err := r.publisher.TaskUpdated(ctx, publisher.TaskUpdated{
		Id:          newVal.UID(),
		Title:       newVal.Title(),
		Description: newVal.Description(),
		AssignedTo:  newVal.AssignedTo().UID,
		CreatedBy:   newVal.CreatedBy().UID,
		Status:      string(newVal.Status()),
	}); err != nil {
		return err
	}

	return nil
}
