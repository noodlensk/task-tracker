package command

import (
	"context"
	"fmt"

	"github.com/noodlensk/task-tracker/internal/accounting/domain/account"
	"github.com/noodlensk/task-tracker/internal/accounting/domain/task"
)

type PayForFinishedTask struct {
	TaskUID string
}

type PayedForFinishedTask struct {
	TaskUID string
	UserUID string
	Amount  float32
}

type PayForFinishedTaskHandler struct {
	taskRepo       task.Repository
	accountRepo    account.Repository
	eventPublisher PayForFinishedTaskEventPublisher
}

func NewPayForFinishedTaskHandler(taskRepo task.Repository, accountRepo account.Repository, eventPublisher PayForFinishedTaskEventPublisher) PayForFinishedTaskHandler {
	return PayForFinishedTaskHandler{taskRepo: taskRepo, accountRepo: accountRepo, eventPublisher: eventPublisher}
}

type PayForFinishedTaskEventPublisher interface {
	PayedForFinishedTask(ctx context.Context, e PayedForFinishedTask)
}

func (h *PayForFinishedTaskHandler) Handle(ctx context.Context, cmd PayForFinishedTask) error {
	t, err := h.taskRepo.Get(ctx, cmd.TaskUID)
	if err != nil {
		return err
	}

	res := &PayedForFinishedTask{
		TaskUID: cmd.TaskUID,
		UserUID: t.AssignedToUserUID(),
	}

	if err := h.accountRepo.Update(ctx, t.AssignedToUserUID(), func(ctx context.Context, a *account.Account) (*account.Account, error) {
		reason := fmt.Sprintf("Task [%s]%s finished", t.UID(), t.Title())
		debitTransaction, err := account.NewDebitTransaction(t.AssignedToUserUID(), t.PriceFinished(), reason)
		if err != nil {
			return nil, err
		}

		if err := a.AddTransaction(debitTransaction); err != nil {
			return nil, err
		}

		res.Amount = debitTransaction.Debit()

		return a, nil
	}); err != nil {
		return err
	}

	h.eventPublisher.PayedForFinishedTask(ctx, PayedForFinishedTask{
		UserUID: t.AssignedToUserUID(),
		TaskUID: cmd.TaskUID,
		Amount:  t.PriceAssigned(),
	})

	return nil
}
