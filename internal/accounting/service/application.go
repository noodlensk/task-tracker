package service

import (
	"github.com/ThreeDotsLabs/watermill/message"

	"github.com/noodlensk/task-tracker/internal/accounting/adapters"
	"github.com/noodlensk/task-tracker/internal/accounting/app"
	"github.com/noodlensk/task-tracker/internal/accounting/app/command"
	"github.com/noodlensk/task-tracker/internal/accounting/app/cud"
	"github.com/noodlensk/task-tracker/internal/accounting/domain/account"
	"github.com/noodlensk/task-tracker/internal/accounting/domain/task"
	"github.com/noodlensk/task-tracker/internal/common/tests"
)

func NewApplication(publisher message.Publisher) (*app.Application, error) {
	taskRepo := adapters.NewTaskInMemoryRepository()
	accountRepo := adapters.NewAccountInMemoryRepository()
	eventPublisher := adapters.NewAsyncEventPublisher(publisher)

	return newApplication(taskRepo, accountRepo, eventPublisher), nil
}

func NewComponentTestApplication() *app.Application {
	pub, err := tests.NewAsyncPublisher()
	if err != nil {
		panic(err)
	}

	taskRepo := adapters.NewTaskInMemoryRepository()
	accountRepo := adapters.NewAccountInMemoryRepository()
	eventPublisher := adapters.NewAsyncEventPublisher(pub)

	return newApplication(taskRepo, accountRepo, eventPublisher)
}

func newApplication(tasksRepo task.Repository, accountRepo account.Repository, eventPublisher command.EventPublisher) *app.Application {
	return &app.Application{
		Commands: app.Commands{
			EstimateTaskPrice:  command.NewEstimateTaskPriceHandler(tasksRepo, eventPublisher),
			PayForFinishedTask: command.NewPayForFinishedTaskHandler(tasksRepo, accountRepo, eventPublisher),
		},
		CUDEvents: app.CUDEvents{
			TaskCreated: cud.NewTaskCreatedHandler(tasksRepo),
		},
	}
}
