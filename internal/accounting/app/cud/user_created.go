package cud

import (
	"context"

	"github.com/noodlensk/task-tracker/internal/accounting/domain/user"
)

type UserCreated struct {
	User user.User
}

type UserCreatedEventHandler struct {
	repo user.Repository
}

func NewUserCreatedEventHandler(repo user.Repository) UserCreatedEventHandler {
	return UserCreatedEventHandler{repo: repo}
}

func (h UserCreatedEventHandler) Handle(ctx context.Context, cmd UserCreated) error {
	return h.repo.StoreUser(ctx, cmd.User)
}
