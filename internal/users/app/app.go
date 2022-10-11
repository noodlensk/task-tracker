package app

import (
	"github.com/noodlensk/task-tracker/internal/users/app/command"
	"github.com/noodlensk/task-tracker/internal/users/app/query"
)

type Application struct {
	Commands Commands
	Queries  Queries
}

type Commands struct {
	CreateUser command.CreateUserHandler
}

type Queries struct {
	AllUsers  query.AllUsersHandler
	AuthLogin query.AuthLoginHandler
}
