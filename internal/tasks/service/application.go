package service

import (
	"github.com/noodlensk/task-tracker/internal/tasks/adapters"
	"github.com/noodlensk/task-tracker/internal/tasks/app"
	"github.com/noodlensk/task-tracker/internal/tasks/app/command"
	"github.com/noodlensk/task-tracker/internal/tasks/app/cud"
	"github.com/noodlensk/task-tracker/internal/tasks/app/query"
	"github.com/noodlensk/task-tracker/internal/tasks/domain/task"
	"github.com/noodlensk/task-tracker/internal/tasks/domain/user"
)

func NewApplication() (*app.Application, error) {
	taskRepo := adapters.NewTaskInMemoryRepository()
	userRepo := adapters.NewUserInMemoryRepository()

	return newApplication(userRepo, taskRepo), nil
}

func NewComponentTestApplication() *app.Application {
	taskRepo := adapters.NewTaskInMemoryRepository()
	userRepo := adapters.NewUserInMemoryRepository()

	return newApplication(userRepo, taskRepo)
}

func newApplication(userRepo user.Repository, taskRepo task.Repository) *app.Application {
	return &app.Application{
		Commands: app.Commands{
			CompleteTask:     command.NewCompleteTaskHandler(taskRepo),
			CreateTask:       command.NewCreateTaskHandler(userRepo, taskRepo),
			ReAssignAllTasks: command.NewReAssignAllTasksHandler(userRepo, taskRepo),
		},
		CUDEvents: app.CUDEvents{
			UserCreated: cud.NewUserCreatedEventHandler(userRepo),
			UserUpdated: cud.NewUserUpdatedEventHandler(userRepo),
		},
		Queries: app.Queries{
			AllTasks:     query.NewAllTasksHandler(taskRepo),
			TasksForUser: query.NewTasksForUserHandler(taskRepo),
		},
	}
}