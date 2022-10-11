package async

import (
	"context"

	"github.com/noodlensk/task-tracker/internal/tasks/app"
	"github.com/noodlensk/task-tracker/internal/tasks/app/cud"
	"github.com/noodlensk/task-tracker/internal/tasks/domain/user"
)

type Server struct {
	app *app.Application
}

func (a *Server) TasksReAssign(ctx context.Context, e TasksReAssign) error {
	// TODO implement me once we will have async for it
	panic("implement me")
}

func (a *Server) UserCreated(ctx context.Context, e UserCreated) error {
	return a.app.CUDEvents.UserCreated.Handle(ctx, cud.UserCreated{User: user.User{
		UID:   e.Id,
		Name:  e.Name,
		Email: e.Email,
		Role:  e.Role,
	}})
}

func (a *Server) UserUpdated(ctx context.Context, e UserUpdated) error {
	return a.app.CUDEvents.UserUpdated.Handle(ctx, cud.UserUpdated{User: user.User{
		UID:   e.Id,
		Name:  e.Name,
		Email: e.Email,
		Role:  e.Role,
	}})
}

func NewAsyncServer(application *app.Application) *Server {
	return &Server{app: application}
}
