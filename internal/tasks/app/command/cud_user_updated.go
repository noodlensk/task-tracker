package command

import (
	"context"
	"github.com/noodlensk/task-tracker/internal/tasks/domain/user"
)

type CUDUserUpdated struct {
	User user.User
}

type CUDUserUpdatedHandler struct {
	repo user.Repository
}

func NewCUDUserUpdatedHandler(repo user.Repository) CUDUserUpdatedHandler {
	return CUDUserUpdatedHandler{repo: repo}
}

func (h CUDUserUpdatedHandler) Handle(ctx context.Context, cmd CUDUserUpdated) error {
	return nil
}
