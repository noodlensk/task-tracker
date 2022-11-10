package query

import (
	"context"

	"github.com/noodlensk/task-tracker/internal/accounting/domain/account"
)

type UserAccountParams struct {
	UserUID string
}

type UserAccount struct {
	Account account.Account
}

type UserAccountHandler struct {
	accountRepo account.Repository
}

func (h *UserAccountHandler) Handle(ctx context.Context, q UserAccountParams) (*UserAccount, error) {
	a, err := h.accountRepo.Get(ctx, q.UserUID)
	if err != nil {
		return nil, err
	}

	return &UserAccount{Account: *a}, nil
}
