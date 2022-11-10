package async

import (
	"context"

	"github.com/noodlensk/task-tracker/internal/tasks/app"
)

type Server struct {
	app *app.Application
}

func (a *Server) ReAssignTasks(ctx context.Context, e ReAssignTasks) error {
	// TODO implement me once we will have async for it
	panic("implement me")
}

func NewAsyncServer(application *app.Application) *Server {
	return &Server{app: application}
}
