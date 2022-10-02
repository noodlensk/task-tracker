package user

import (
	"context"
)

type Repository interface {
	CreateUser(ctx context.Context, u *User) error
	GetAllUsers(ctx context.Context, limit, offset uint) ([]*User, error)
	GetUserByEmail(ctx context.Context, email string) (*User, error)

	UpdateUser(
		ctx context.Context,
		userUID string,
		updateFn func(ctx context.Context, u *User) (*User, error),
	) error
}
