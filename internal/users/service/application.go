package service

import (
	"github.com/noodlensk/task-tracker/internal/common/tests"
	"github.com/noodlensk/task-tracker/internal/users/adapters"
	"github.com/noodlensk/task-tracker/internal/users/app"
	"github.com/noodlensk/task-tracker/internal/users/app/command"
	"github.com/noodlensk/task-tracker/internal/users/app/query"
	"github.com/noodlensk/task-tracker/internal/users/domain/user"
	"go.uber.org/zap"
)

func NewApplication(logger *zap.SugaredLogger) (*app.Application, error) {
	userRepo := adapters.NewUserInMemoryRepository()

	return newApplication(logger, userRepo), nil
}

func NewComponentTestApplication() *app.Application {
	userRepo := adapters.NewUserInMemoryRepository()
	return newApplication(tests.NewLogger(), userRepo)
}

func newApplication(logger *zap.SugaredLogger, userRepo user.Repository) *app.Application {
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
