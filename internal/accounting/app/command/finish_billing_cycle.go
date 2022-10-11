package command

import (
	"context"

	"github.com/noodlensk/task-tracker/internal/accounting/domain/account"
)

type FinishBillingCycle struct{}

type BillingCycleFinished struct {
	UserUID string
}

type FinishBillingCycleHandler struct {
	accountRepo account.Repository
}

func (h *FinishBillingCycleHandler) Handle(ctx context.Context, cmd FinishBillingCycle) ([]*BillingCycleFinished, error) {
	var res []*BillingCycleFinished

	accounts, err := h.accountRepo.List(ctx)
	if err != nil {
		return nil, err
	}

	for _, a := range accounts {
		if err := h.accountRepo.Update(ctx, a.UserUID(), func(ctx context.Context, a *account.Account) (*account.Account, error) {
			if err := a.CloseBillingCycle(); err != nil {
				return nil, err
			}

			return a, nil
		}); err != nil {
			return nil, err
		}

		res = append(res, &BillingCycleFinished{
			UserUID: a.UserUID(),
		})
	}

	return res, nil
}
