package app

import (
	"github.com/noodlensk/task-tracker/internal/tasks/app/command"
	"github.com/noodlensk/task-tracker/internal/tasks/app/cud"
	"github.com/noodlensk/task-tracker/internal/tasks/app/query"
)

type Application struct {
	Commands  Commands
	Queries   Queries
	CUDEvents CUDEvents
}

type Commands struct {
	CreateTask       command.CreateTaskHandler
	CompleteTask     command.CompleteTaskHandler
	ReAssignAllTasks command.ReAssignAllTasksHandler
}

type CUDEvents struct {
	UserCreated cud.UserCreatedEventHandler
	UserUpdated cud.UserUpdatedEventHandler
}

type Queries struct {
	AllTasks     query.AllTasksHandler
	TasksForUser query.TasksForUserHandler
}
