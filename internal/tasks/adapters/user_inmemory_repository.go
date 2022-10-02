package adapters

import (
	"context"
	"github.com/noodlensk/task-tracker/internal/tasks/domain/user"
)

type UserInMemoryRepository struct {
}

func NewUserInMemoryRepository() *UserInMemoryRepository {
	return &UserInMemoryRepository{}
}

func (r UserInMemoryRepository) GetUsers(ctx context.Context, role string) ([]*user.User, error) {
	//TODO implement me
	panic("implement me")
}
