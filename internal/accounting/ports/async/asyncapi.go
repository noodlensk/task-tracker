package async

import (
	"context"

	"github.com/noodlensk/task-tracker/internal/accounting/app"
	"github.com/noodlensk/task-tracker/internal/accounting/app/command"
)

type Server struct {
	app *app.Application
}

func (s Server) TaskCreated(ctx context.Context, e TaskCreated) error {
	if err := s.app.Commands.EstimateTaskPrice.Handle(ctx, command.EstimateTaskPrice{
		TaskUID: e.PublicId,
	}); err != nil {
		return err
	}

	return nil
}

func (s Server) TaskAssigned(ctx context.Context, e TaskAssigned) error {
	return s.app.Commands.ChargeForAssignedTask.Handle(ctx, command.ChargeForAssignedTask{
		TaskUID: e.PublicId,
		UserUID: e.UserId,
	})
}

func (s Server) TaskCompleted(ctx context.Context, e TaskCompleted) error {
	return s.app.Commands.PayForFinishedTask.Handle(ctx, command.PayForFinishedTask{
		TaskUID: e.PublicId,
	})
}

func NewAsyncServer(application *app.Application) *Server {
	return &Server{app: application}
}
