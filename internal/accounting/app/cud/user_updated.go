package cud

import (
	"context"

	"github.com/noodlensk/task-tracker/internal/accounting/domain/user"
)

type UserUpdated struct {
	User user.User
}

type UserUpdatedEventHandler struct {
	repo user.Repository
}

func NewUserUpdatedEventHandler(repo user.Repository) UserUpdatedEventHandler {
	return UserUpdatedEventHandler{repo: repo}
}

func (h UserUpdatedEventHandler) Handle(ctx context.Context, cmd UserUpdated) error {
	return h.repo.StoreUser(ctx, cmd.User)
}
