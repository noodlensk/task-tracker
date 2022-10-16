package adapters

import (
	"context"
	"errors"
	"sync"

	"github.com/noodlensk/task-tracker/internal/users/data/publisher"
	"github.com/noodlensk/task-tracker/internal/users/domain/user"
)

type UserInMemoryRepository struct {
	users []*user.User

	publisher *publisher.PublisherClient

	lock sync.RWMutex
}

func NewUserInMemoryRepository(publisher *publisher.PublisherClient) *UserInMemoryRepository {
	adminUser, _ := user.NewUser("admin", "admin@admin.com", user.RoleAdmin, "admin")

	return &UserInMemoryRepository{
		users:     []*user.User{adminUser},
		publisher: publisher,
	}
}

func (r *UserInMemoryRepository) CreateUser(ctx context.Context, u *user.User) error {
	r.lock.Lock()
	defer r.lock.Unlock()

	r.users = append(r.users, u)

	if err := r.publisher.UserCreated(ctx, publisher.UserCreated{
		Id:    u.UID(),
		Name:  u.Name(),
		Email: u.Email(),
		Role:  u.Role().String(),
	}); err != nil {
		return err
	}

	return nil
}

func (r *UserInMemoryRepository) GetAllUsers(ctx context.Context, limit, offset uint) ([]*user.User, error) {
	r.lock.RLock()
	defer r.lock.RUnlock()

	var res []*user.User

	if offset > uint(len(r.users)) {
		return nil, nil
	}

	for _, v := range r.users[offset:] {
		if uint(len(res)) >= limit {
			break
		}

		res = append(res, v)
	}

	return res, nil
}

func (r *UserInMemoryRepository) UpdateUser(ctx context.Context, userUID string, updateFn func(ctx context.Context, u *user.User) (*user.User, error)) error {
	r.lock.Lock()
	defer r.lock.Unlock()

	var (
		userIdx   int
		userFound bool
	)

	for i, v := range r.users {
		if v.UID() == userUID {
			userIdx = i
			userFound = true

			break
		}
	}

	if !userFound {
		return errors.New("not found")
	}

	newVal, err := updateFn(ctx, r.users[userIdx])
	if err != nil {
		return err
	}

	r.users[userIdx] = newVal

	if err := r.publisher.UserUpdated(ctx, publisher.UserUpdated{
		Id:    newVal.UID(),
		Name:  newVal.Name(),
		Email: newVal.Email(),
		Role:  newVal.Role().String(),
	}); err != nil {
		return err
	}

	return nil
}

func (r *UserInMemoryRepository) GetUserByEmail(ctx context.Context, email string) (*user.User, error) {
	r.lock.RLock()
	defer r.lock.RUnlock()

	for _, v := range r.users {
		if v.Email() == email {
			return v, nil
		}
	}

	return nil, errors.New("not found")
}
