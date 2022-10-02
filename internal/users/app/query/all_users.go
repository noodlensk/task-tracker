package query

import (
	"context"

	"github.com/noodlensk/task-tracker/internal/users/domain/user"
)

type AllUsers struct {
	Limit  int
	Offset int
}

type AllUsersHandler struct {
	repo user.Repository
}

func NewAllUsersHandler(repo user.Repository) AllUsersHandler {
	return AllUsersHandler{repo: repo}
}

func (h AllUsersHandler) Handle(ctx context.Context, q AllUsers) ([]*user.User, error) {
	return h.repo.GetAllUsers(ctx, uint(q.Limit), uint(q.Offset))
}
