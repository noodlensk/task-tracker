package command

import (
	"context"

	"github.com/ThreeDotsLabs/watermill/message"

	decorator "github.com/noodlensk/task-tracker/internal/common/decorators"
	"github.com/noodlensk/task-tracker/internal/users/domain/user"
)

type CreateUser struct {
	User user.User
}

type UserCreated struct {
	User user.User
}

type CreateUserHandler decorator.CommandHandler[CreateUser, *UserCreated]

type createUserHandler struct {
	repo user.Repository
}

func NewCreateUserHandler(repo user.Repository, publisher message.Publisher, topic string) CreateUserHandler {
	return decorator.ApplyCommandDecorators[CreateUser, *UserCreated](createUserHandler{repo: repo}, publisher, topic)
}

func (h createUserHandler) Handle(ctx context.Context, cmd CreateUser) (*UserCreated, error) {
	if err := h.repo.CreateUser(ctx, &cmd.User); err != nil {
		return nil, err
	}

	return &UserCreated{User: cmd.User}, nil
}
