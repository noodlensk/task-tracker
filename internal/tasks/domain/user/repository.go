package user

import "context"

type Repository interface {
	GetUsers(ctx context.Context, role string) ([]*User, error)
	StoreUser(ctx context.Context, u User) error
}
