package command

import (
	"context"
	"github.com/noodlensk/task-tracker/internal/tasks/domain/user"
)

type CUDUserCreated struct {
	User user.User
}

type CUDUserCreatedHandler struct {
	repo user.Repository
}

func NewCUDUserCreatedHandler(repo user.Repository) CUDUserCreatedHandler {
	return CUDUserCreatedHandler{repo: repo}
}

func (h CUDUserCreatedHandler) Handle(ctx context.Context, cmd CUDUserCreated) error {
	return nil
}
