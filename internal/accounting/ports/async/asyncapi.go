package async

import (
	"context"

	"github.com/noodlensk/task-tracker/internal/accounting/app"
	"github.com/noodlensk/task-tracker/internal/accounting/app/command"
	"github.com/noodlensk/task-tracker/internal/accounting/app/cud"
	"github.com/noodlensk/task-tracker/internal/accounting/domain/task"
)

type Server struct {
	app *app.Application
}

func (s Server) TaskCreated(ctx context.Context, e TaskCreated) error {
	if err := s.app.CUDEvents.TaskCreated.Handle(ctx, cud.TaskCreated{
		Task: task.NewTaskFromCUD(e.Id, e.Title, e.Title, e.AssignedTo),
	}); err != nil {
		return err
	}

	if err := s.app.Commands.EstimateTaskPrice.Handle(ctx, command.EstimateTaskPrice{
		TaskUID: e.Id,
	}); err != nil {
		return err
	}

	return nil
}

func (s Server) UserCreated(ctx context.Context, e UserCreated) error {
	return nil
}

func (s Server) UserUpdated(ctx context.Context, e UserUpdated) error {
	return nil
}

func (s Server) TaskAssigned(ctx context.Context, e TaskAssigned) error {
	if err := s.app.CUDEvents.TaskUpdated.Handle(ctx, cud.TaskUpdated{
		Task: task.NewTaskFromCUD(e.Id, e.Title, e.AssignedTo, e.Status),
	}); err != nil {
		return err
	}

	return s.app.Commands.ChargeForAssignedTask.Handle(ctx, command.ChargeForAssignedTask{
		TaskUID: e.Id,
		UserUID: e.AssignedTo,
	})
}

func (s Server) TaskCompleted(ctx context.Context, e TaskCompleted) error {
	if err := s.app.CUDEvents.TaskUpdated.Handle(ctx, cud.TaskUpdated{
		Task: task.NewTaskFromCUD(e.Id, e.Title, e.AssignedTo, e.Status),
	}); err != nil {
		return err
	}

	return s.app.Commands.PayForFinishedTask.Handle(ctx, command.PayForFinishedTask{
		TaskUID: e.Id,
	})
}

func NewAsyncServer(application *app.Application) *Server {
	return &Server{app: application}
}
