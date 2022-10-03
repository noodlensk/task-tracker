package service

import (
	"github.com/ThreeDotsLabs/watermill/message"

	"github.com/noodlensk/task-tracker/internal/users/adapters"
	"github.com/noodlensk/task-tracker/internal/users/app"
	"github.com/noodlensk/task-tracker/internal/users/app/command"
	"github.com/noodlensk/task-tracker/internal/users/app/query"
	"github.com/noodlensk/task-tracker/internal/users/domain/user"
)

func NewApplication(publisher message.Publisher) (*app.Application, error) {
	userRepo := adapters.NewUserInMemoryRepository()

	return newApplication(userRepo, publisher), nil
}

func NewComponentTestApplication() *app.Application {
	userRepo := adapters.NewUserInMemoryRepository()

	return newApplication(userRepo, nil)
}

func newApplication(userRepo user.Repository, publisher message.Publisher) *app.Application {
	return &app.Application{
		Commands: app.Commands{
			CreateUser: command.NewCreateUserHandler(userRepo, publisher, "user.created"),
		},
		Queries: app.Queries{
			AllUsers:  query.NewAllUsersHandler(userRepo),
			AuthLogin: query.NewAuthLoginHandler(userRepo, "secret"),
		},
	}
}
