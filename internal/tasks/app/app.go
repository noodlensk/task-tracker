package app

import (
	"github.com/noodlensk/task-tracker/internal/tasks/app/command"
	"github.com/noodlensk/task-tracker/internal/tasks/app/query"
)

type Application struct {
	Commands Commands
	Queries  Queries
}

type Commands struct {
	CreateTask       command.CreateTaskHandler
	CompleteTask     command.CompleteTaskHandler
	ReAssignAllTasks command.ReAssignAllTasksHandler
}

type Queries struct {
	AllTasks     query.AllTasksHandler
	TasksForUser query.TasksForUserHandler
}
