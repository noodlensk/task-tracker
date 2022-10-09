package adapters

import (
	"context"
	"sync"

	"github.com/noodlensk/task-tracker/internal/tasks/domain/user"
)

type UserInMemoryRepository struct {
	users map[string]*user.User

	lock sync.RWMutex
}

func NewUserInMemoryRepository() *UserInMemoryRepository {
	return &UserInMemoryRepository{users: map[string]*user.User{}}
}

func (r *UserInMemoryRepository) StoreUser(ctx context.Context, u user.User) error {
	r.lock.Lock()
	defer r.lock.Unlock()

	r.users[u.UID] = &u

	return nil
}

func (r *UserInMemoryRepository) GetUsers(ctx context.Context, role string) ([]*user.User, error) {
	r.lock.RLock()
	defer r.lock.RUnlock()

	var res []*user.User

	for _, u := range r.users {
		if u.Role == role {
			res = append(res, u)
		}
	}

	return res, nil
}
