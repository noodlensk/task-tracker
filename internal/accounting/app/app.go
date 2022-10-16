package app

import (
	"github.com/noodlensk/task-tracker/internal/accounting/app/command"
	"github.com/noodlensk/task-tracker/internal/accounting/app/cud"
)

type Application struct {
	Commands  Commands
	Queries   Queries
	CUDEvents CUDEvents
}

type Commands struct {
	EstimateTaskPrice     command.EstimateTaskPriceHandler
	ChargeForAssignedTask command.ChargeForAssignedTaskHandler
	PayForFinishedTask    command.PayForFinishedTaskHandler
}

type CUDEvents struct {
	TaskCreated cud.TaskCreatedEventHandler
	TaskUpdated cud.TaskUpdatedEventHandler
	UserCreated cud.UserCreatedEventHandler
	UserUpdated cud.UserUpdatedEventHandler
}

type Queries struct{}
