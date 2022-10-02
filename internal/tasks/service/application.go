package service

import (
	"go.uber.org/zap"

	"github.com/noodlensk/task-tracker/internal/tasks/adapters"
	"github.com/noodlensk/task-tracker/internal/tasks/app"
	"github.com/noodlensk/task-tracker/internal/tasks/app/command"
	"github.com/noodlensk/task-tracker/internal/tasks/app/query"
	"github.com/noodlensk/task-tracker/internal/tasks/domain/task"
	"github.com/noodlensk/task-tracker/internal/tasks/domain/user"
)

func NewApplication(logger *zap.SugaredLogger) (*app.Application, error) {
	taskRepo := adapters.NewTaskInMemoryRepository()
	userRepo := adapters.NewUserInMemoryRepository()

	return newApplication(logger, userRepo, taskRepo), nil
}

//func NewComponentTestApplication() *app.Application {
//	return newApplication(tests.NewLogger())
//}

func newApplication(logger *zap.SugaredLogger, userRepo user.Repository, taskRepo task.Repository) *app.Application {
	return &app.Application{
		Commands: app.Commands{
			CompleteTask:     command.NewCompleteTaskHandler(taskRepo),
			CreateTask:       command.NewCreateTaskHandler(taskRepo),
			ReAssignAllTasks: command.NewReAssignAllTasksHandler(userRepo, taskRepo),
		},
		Queries: app.Queries{
			AllTasks:     query.NewAllTasksHandler(taskRepo),
			TasksForUser: query.NewTasksForUserHandler(taskRepo),
		},
	}
}
