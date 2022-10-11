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
	TaskCreated cud.TaskCreatedHandler
	TaskUpdated cud.TaskUpdatedHandler
}

type Queries struct{}
