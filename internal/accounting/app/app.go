package app

import (
	"github.com/noodlensk/task-tracker/internal/accounting/app/command"
)

type Application struct {
	Commands Commands
	Queries  Queries
}

type Commands struct {
	EstimateTaskPrice     command.EstimateTaskPriceHandler
	ChargeForAssignedTask command.ChargeForAssignedTaskHandler
	PayForFinishedTask    command.PayForFinishedTaskHandler
}

type Queries struct{}
