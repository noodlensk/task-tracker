package subscriber

import (
	"context"

	"github.com/noodlensk/task-tracker/internal/accounting/domain/task"
	"github.com/noodlensk/task-tracker/internal/accounting/domain/user"
)

type TasksRepository interface {
	Create(ctx context.Context, task *task.Task) error
	Update(ctx context.Context, uid string, updateFn func(ctx context.Context, task *task.Task) (*task.Task, error)) error
}

type UsersRepository interface {
	StoreUser(ctx context.Context, u user.User) error
}

type Server struct {
	tasksRepo TasksRepository
	usersRepo UsersRepository
}

func (s Server) TaskCreated(ctx context.Context, e TaskCreated) error {
	t := task.NewTaskFromCUD(e.Id, e.Title, e.Title, e.AssignedTo)

	if err := s.tasksRepo.Create(ctx, &t); err != nil {
		return err
	}

	return nil
}

func (s Server) TaskUpdated(ctx context.Context, e TaskUpdated) error {
	t := task.NewTaskFromCUD(e.Id, e.Title, e.Title, e.AssignedTo)

	if err := s.tasksRepo.Update(ctx, t.UID(), func(ctx context.Context, t *task.Task) (*task.Task, error) {
		t.SetTitle(e.Title)
		t.SetStatus(e.Status)

		return t, nil
	}); err != nil {
		return err
	}

	return nil
}

func (s Server) UserCreated(ctx context.Context, e UserCreated) error {
	return s.usersRepo.StoreUser(ctx, user.User{
		UID:   e.Id,
		Name:  e.Name,
		Email: e.Email,
		Role:  e.Role,
	})
}

func (s Server) UserUpdated(ctx context.Context, e UserUpdated) error {
	return s.usersRepo.StoreUser(ctx, user.User{
		UID:   e.Id,
		Name:  e.Name,
		Email: e.Email,
		Role:  e.Role,
	})
}

func NewAsyncServer(tasksRepo TasksRepository, usersRepo UsersRepository) *Server {
	return &Server{tasksRepo: tasksRepo, usersRepo: usersRepo}
}
