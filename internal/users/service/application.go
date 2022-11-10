package service

import (
	"github.com/ThreeDotsLabs/watermill/message"

	"github.com/noodlensk/task-tracker/internal/common/tests"
	"github.com/noodlensk/task-tracker/internal/users/adapters"
	"github.com/noodlensk/task-tracker/internal/users/app"
	"github.com/noodlensk/task-tracker/internal/users/app/command"
	"github.com/noodlensk/task-tracker/internal/users/app/query"
	"github.com/noodlensk/task-tracker/internal/users/data/publisher"
	"github.com/noodlensk/task-tracker/internal/users/domain/user"
)

func NewApplication(pub message.Publisher) (*app.Application, error) {
	dataPublisher := publisher.NewPublisherClient(pub)

	userRepo := adapters.NewUserInMemoryRepository(dataPublisher)

	return newApplication(userRepo), nil
}

func NewComponentTestApplication() *app.Application {
	pub, err := tests.NewAsyncPublisher()
	if err != nil {
		panic(err)
	}

	dataPublisher := publisher.NewPublisherClient(pub)
	userRepo := adapters.NewUserInMemoryRepository(dataPublisher)

	return newApplication(userRepo)
}

func newApplication(userRepo user.Repository) *app.Application {
	return &app.Application{
		Commands: app.Commands{
			CreateUser: command.NewCreateUserHandler(userRepo),
		},
		Queries: app.Queries{
			AllUsers:  query.NewAllUsersHandler(userRepo),
			AuthLogin: query.NewAuthLoginHandler(userRepo, "secret"),
		},
	}
}
