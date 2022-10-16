package command

import (
	"context"

	"github.com/noodlensk/task-tracker/internal/users/domain/user"
)

type CreateUser struct {
	User user.User
}

type UserCreated struct {
	User user.User
}

type CreateUserHandler struct {
	repo user.Repository
}

func NewCreateUserHandler(repo user.Repository) CreateUserHandler {
	return CreateUserHandler{repo: repo}
}

func (h CreateUserHandler) Handle(ctx context.Context, cmd CreateUser) (*UserCreated, error) {
	if err := h.repo.CreateUser(ctx, &cmd.User); err != nil {
		return nil, err
	}

	return &UserCreated{User: cmd.User}, nil
}
