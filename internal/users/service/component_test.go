package service

import (
	"context"
	"log"
	"net/http"
	"os"
	"testing"

	"github.com/go-chi/chi/v5"
	"github.com/google/uuid"
	"github.com/noodlensk/task-tracker/internal/common/clients/users"
	"github.com/noodlensk/task-tracker/internal/common/server"
	"github.com/noodlensk/task-tracker/internal/common/tests"
	"github.com/noodlensk/task-tracker/internal/users/ports"
	"github.com/pkg/errors"
	"github.com/stretchr/testify/require"
)

func TestGetUsers(t *testing.T) {
	t.Parallel()

	token := tests.FakeAdminJWT(t, uuid.New().String())
	client := tests.NewUsersHTTPClient(t, token)

	usersList := client.GetUsers(t)

	require.True(t, len(usersList) > 0)
}

func TestCreateUser(t *testing.T) {
	t.Parallel()

	token := tests.FakeAdminJWT(t, uuid.New().String())
	client := tests.NewUsersHTTPClient(t, token)

	userToCreate := users.CreateUserRequest{
		Email:    "myemail@google.com",
		Name:     "my user",
		Password: "my password",
		Role:     users.CreateUserRequestRoleManager,
	}

	client.CreateUser(t, userToCreate)

	userList := client.GetUsers(t)
	userFound := false

	for _, u := range userList {
		if u.Email != nil && *u.Email == userToCreate.Email {
			userFound = true

			require.Equal(t, *u.Name, userToCreate.Name)
			require.Equal(t, string(u.Role), string(userToCreate.Role))
		}
	}

	require.True(t, userFound)
}

func TestAuthLogin(t *testing.T) {
	t.Parallel()

	token := tests.FakeAdminJWT(t, uuid.New().String())
	client := tests.NewUsersHTTPClient(t, token)

	userToCreate := users.CreateUserRequest{
		Email:    "myloginemail@google.com",
		Name:     "my user",
		Password: "my password",
		Role:     users.CreateUserRequestRoleManager,
	}

	client.CreateUser(t, userToCreate)

	newToken := client.AuthLogin(t, userToCreate.Email, userToCreate.Password)

	require.NotEmpty(t, newToken)
}

func startService() error {
	app := NewComponentTestApplication()
	ctx := context.Background()

	httpAddr := "127.0.0.1:8081"

	go server.RunHTTPServerOnAddr(ctx, httpAddr, tests.NewLogger(), func(router chi.Router) http.Handler {
		return ports.HandlerFromMux(ports.NewHTTPServer(app), router)
	})

	ok := tests.WaitForPort(httpAddr)
	if !ok {
		return errors.Errorf("Timed out waiting for trainings HTTP to come up")
	}

	return nil
}

func TestMain(m *testing.M) {
	if err := startService(); err != nil {
		log.Printf("Failed to start service: %q\n", err.Error())
		os.Exit(1)
	}

	os.Exit(m.Run())
}
