package subscriber

import (
	"context"

	"github.com/noodlensk/task-tracker/internal/tasks/domain/user"
)

type UsersRepository interface {
	StoreUser(ctx context.Context, u user.User) error
}

type Server struct {
	usersRepo UsersRepository
}

func (a *Server) UserCreated(ctx context.Context, e UserCreated) error {
	return a.usersRepo.StoreUser(ctx, user.User{
		UID:   e.Id,
		Name:  e.Name,
		Email: e.Email,
		Role:  e.Role,
	})
}

func (a *Server) UserUpdated(ctx context.Context, e UserUpdated) error {
	return a.usersRepo.StoreUser(ctx, user.User{
		UID:   e.Id,
		Name:  e.Name,
		Email: e.Email,
		Role:  e.Role,
	})
}

func NewAsyncServer(usersRepo UsersRepository) *Server {
	return &Server{usersRepo: usersRepo}
}
