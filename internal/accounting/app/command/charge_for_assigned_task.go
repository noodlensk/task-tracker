package command

import (
	"context"
	"fmt"

	"github.com/noodlensk/task-tracker/internal/accounting/domain/account"
	"github.com/noodlensk/task-tracker/internal/accounting/domain/task"
)

type ChargeForAssignedTask struct {
	TaskUID string
	UserUID string
}

type ChargedForAssignedTask struct {
	TaskUID string
	UserUID string
	Amount  float32
}

type ChargeForAssignedTaskEventPublisher interface {
	UserChargedForAssignedTask(ctx context.Context, e ChargedForAssignedTask)
}

type ChargeForAssignedTaskHandler struct {
	taskRepo       task.Repository
	accountRepo    account.Repository
	eventPublisher ChargeForAssignedTaskEventPublisher
}

func NewChargeForAssignedTaskHandler(taskRepo task.Repository, accountRepo account.Repository, eventPublisher ChargeForAssignedTaskEventPublisher) ChargeForAssignedTaskHandler {
	return ChargeForAssignedTaskHandler{taskRepo: taskRepo, accountRepo: accountRepo, eventPublisher: eventPublisher}
}

func (h ChargeForAssignedTaskHandler) Handle(ctx context.Context, cmd ChargeForAssignedTask) error {
	t, err := h.taskRepo.Get(ctx, cmd.TaskUID)
	if err != nil {
		return err
	}

	if err := h.accountRepo.Update(ctx, cmd.UserUID, func(ctx context.Context, a *account.Account) (*account.Account, error) {
		reason := fmt.Sprintf("Task [%s]%s assigned", t.UID(), t.Title())
		creditTransaction, err := account.NewCreditTransaction(cmd.UserUID, t.PriceAssigned(), reason)
		if err != nil {
			return nil, err
		}

		if err := a.AddTransaction(creditTransaction); err != nil {
			return nil, err
		}

		return a, nil
	}); err != nil {
		return err
	}

	h.eventPublisher.UserChargedForAssignedTask(ctx, ChargedForAssignedTask{
		UserUID: cmd.UserUID,
		TaskUID: cmd.TaskUID,
		Amount:  t.PriceAssigned(),
	})

	return nil
}
