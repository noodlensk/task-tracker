package tests

import (
	"context"
	"fmt"
	"net/http"
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/noodlensk/task-tracker/internal/common/clients/users"
)

func authorizationBearer(token string) func(context.Context, *http.Request) error {
	return func(ctx context.Context, req *http.Request) error {
		req.Header.Add("Authorization", fmt.Sprintf("Bearer %s", token))

		return nil
	}
}

func NewUsersHTTPClient(t *testing.T, token string) UsersHTTPClient {
	t.Helper()

	addr := "localhost:8081"

	t.Log("Trying users http:", addr)
	ok := WaitForPort(addr)
	require.True(t, ok, "Users HTTP timed out")

	url := fmt.Sprintf("http://%v/api", addr)

	client, err := users.NewClientWithResponses(
		url,
		users.WithRequestEditorFn(authorizationBearer(token)),
	)
	require.NoError(t, err)

	return UsersHTTPClient{client: client}
}

type UsersHTTPClient struct {
	client *users.ClientWithResponses
}

func (c UsersHTTPClient) GetUsers(t *testing.T) []users.User {
	t.Helper()

	resp, err := c.client.GetUsersWithResponse(context.Background(), &users.GetUsersParams{Offset: 0, Limit: 1000})

	require.NoError(t, err)
	require.Equal(t, http.StatusOK, resp.StatusCode())

	usersList := *resp.JSON200

	return usersList.Users
}

func (c UsersHTTPClient) CreateUser(t *testing.T, u users.CreateUserRequest) {
	t.Helper()

	resp, err := c.client.CreateUser(context.Background(), u)

	require.NoError(t, err)
	require.NoError(t, resp.Body.Close())

	require.Equal(t, http.StatusOK, resp.StatusCode)
}

func (c UsersHTTPClient) AuthLogin(t *testing.T, email, password string) string {
	t.Helper()

	resp, err := c.client.AuthLoginWithResponse(context.Background(), users.AuthLoginRequest{
		Email:    email,
		Password: password,
	})

	require.NoError(t, err)
	require.Equal(t, http.StatusOK, resp.StatusCode())

	respWithToken := *resp.JSON200

	return respWithToken.Token
}
