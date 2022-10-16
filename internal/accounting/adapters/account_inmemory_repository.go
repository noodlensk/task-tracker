package adapters

import (
	"context"
	"errors"
	"sync"

	"github.com/noodlensk/task-tracker/internal/accounting/domain/account"
)

type AccountInMemoryRepository struct {
	accounts map[string]*account.Account

	lock sync.RWMutex
}

func (r *AccountInMemoryRepository) Create(ctx context.Context, a *account.Account) error {
	r.lock.Lock()
	defer r.lock.Unlock()

	if _, ok := r.accounts[a.UserUID()]; ok {
		return errors.New("already exist")
	}

	r.accounts[a.UserUID()] = a

	return nil
}

func (r *AccountInMemoryRepository) List(ctx context.Context) ([]*account.Account, error) {
	r.lock.RLock()
	defer r.lock.RUnlock()

	var res []*account.Account

	for _, a := range r.accounts {
		res = append(res, a)
	}

	return res, nil
}

func (r *AccountInMemoryRepository) Get(ctx context.Context, uid string) (*account.Account, error) {
	r.lock.RLock()
	defer r.lock.RUnlock()

	v, ok := r.accounts[uid]
	if !ok {
		return nil, errors.New("not found")
	}

	return v, nil
}

func (r *AccountInMemoryRepository) Update(ctx context.Context, uid string, updateFn func(ctx context.Context, a *account.Account) (*account.Account, error)) error {
	r.lock.Lock()
	defer r.lock.Unlock()

	v, ok := r.accounts[uid]
	if !ok {
		var err error

		v, err = account.New(uid)
		if err != nil {
			return err
		}
	}

	newVal, err := updateFn(ctx, v)
	if err != nil {
		return err
	}

	r.accounts[uid] = newVal

	return nil
}

func NewAccountInMemoryRepository() *AccountInMemoryRepository {
	return &AccountInMemoryRepository{
		accounts: map[string]*account.Account{},
	}
}
