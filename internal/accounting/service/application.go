package service

import (
	"github.com/ThreeDotsLabs/watermill/message"
	"go.uber.org/zap"

	"github.com/noodlensk/task-tracker/internal/accounting/adapters"
	"github.com/noodlensk/task-tracker/internal/accounting/app"
	"github.com/noodlensk/task-tracker/internal/accounting/app/command"
	"github.com/noodlensk/task-tracker/internal/accounting/app/cud"
	"github.com/noodlensk/task-tracker/internal/accounting/domain/account"
	"github.com/noodlensk/task-tracker/internal/accounting/domain/task"
	"github.com/noodlensk/task-tracker/internal/accounting/domain/user"
	"github.com/noodlensk/task-tracker/internal/common/logs"
	"github.com/noodlensk/task-tracker/internal/common/tests"
)

func NewApplication(publisher message.Publisher, logger *zap.SugaredLogger) (*app.Application, error) {
	usersRepo := adapters.NewUserInMemoryRepository()
	taskRepo := adapters.NewTaskInMemoryRepository()
	accountRepo := adapters.NewAccountInMemoryRepository()
	eventPublisher := adapters.NewAsyncEventPublisher(publisher, logger.With("component", "event-publisher"))

	return newApplication(usersRepo, taskRepo, accountRepo, eventPublisher), nil
}

func NewComponentTestApplication() *app.Application {
	pub, err := tests.NewAsyncPublisher()
	if err != nil {
		panic(err)
	}

	usersRepo := adapters.NewUserInMemoryRepository()
	taskRepo := adapters.NewTaskInMemoryRepository()
	accountRepo := adapters.NewAccountInMemoryRepository()
	eventPublisher := adapters.NewAsyncEventPublisher(pub, logs.NewLogger()) // TODO: replace with nop logger

	return newApplication(usersRepo, taskRepo, accountRepo, eventPublisher)
}

func newApplication(usersRepo user.Repository, tasksRepo task.Repository, accountRepo account.Repository, eventPublisher command.EventPublisher) *app.Application {
	return &app.Application{
		Commands: app.Commands{
			EstimateTaskPrice:     command.NewEstimateTaskPriceHandler(tasksRepo, eventPublisher),
			PayForFinishedTask:    command.NewPayForFinishedTaskHandler(tasksRepo, accountRepo, eventPublisher),
			ChargeForAssignedTask: command.NewChargeForAssignedTaskHandler(tasksRepo, accountRepo, eventPublisher),
		},
		CUDEvents: app.CUDEvents{
			TaskCreated: cud.NewTaskCreatedEventHandler(tasksRepo),
			TaskUpdated: cud.NewTaskUpdatedEventHandler(tasksRepo),
			UserCreated: cud.NewUserCreatedEventHandler(usersRepo),
			UserUpdated: cud.NewUserUpdatedEventHandler(usersRepo),
		},
	}
}
