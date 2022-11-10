package account

import "context"

type Repository interface {
	Create(ctx context.Context, a *Account) error
	List(ctx context.Context) ([]*Account, error)
	Get(ctx context.Context, uid string) (*Account, error)
	Update(ctx context.Context, uid string, updateFn func(ctx context.Context, a *Account) (*Account, error)) error
}
