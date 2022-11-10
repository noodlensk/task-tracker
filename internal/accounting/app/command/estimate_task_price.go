package command

import (
	"context"
	"math/rand"

	"github.com/noodlensk/task-tracker/internal/accounting/domain/task"
)

type EstimateTaskPrice struct {
	TaskUID string
}

type TaskPriceEstimated struct {
	TaskUID       string
	PriceAssigned float32
	PriceFinished float32
}

type TaskPriceEstimatedEventPublisher interface {
	TaskPriceEstimated(ctx context.Context, e TaskPriceEstimated)
}

type EstimateTaskPriceHandler struct {
	tasksRepo      task.Repository
	eventPublisher TaskPriceEstimatedEventPublisher
}

func NewEstimateTaskPriceHandler(tasksRepo task.Repository, eventPublisher TaskPriceEstimatedEventPublisher) EstimateTaskPriceHandler {
	return EstimateTaskPriceHandler{tasksRepo: tasksRepo, eventPublisher: eventPublisher}
}

//nolint:gosec
func (h *EstimateTaskPriceHandler) Handle(ctx context.Context, cmd EstimateTaskPrice) error {
	priceAssigned := float32(rand.Intn(10) + 10) // rand(10..20)
	priceFinished := float32(rand.Intn(20) + 20) // rand(20..40)

	if err := h.tasksRepo.Update(ctx, cmd.TaskUID, func(ctx context.Context, t *task.Task) (*task.Task, error) {
		if err := t.SetPriceAssigned(priceAssigned); err != nil {
			return nil, err
		}

		if err := t.SetPriceFinished(priceFinished); err != nil {
			return nil, err
		}

		return t, nil
	}); err != nil {
		return err
	}

	h.eventPublisher.TaskPriceEstimated(ctx, TaskPriceEstimated{
		TaskUID:       cmd.TaskUID,
		PriceAssigned: priceAssigned,
		PriceFinished: priceFinished,
	})

	return nil
}
